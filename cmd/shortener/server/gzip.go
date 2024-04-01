package server

import (
	"compress/gzip"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// compressWriter реализует интерфейс http.ResponseWriter и позволяет прозрачно для сервера
type compressWriter struct {
	writer    http.ResponseWriter
	zipWriter *gzip.Writer
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		writer:    w,
		zipWriter: gzip.NewWriter(w),
	}
}

func (c *compressWriter) Header() http.Header {
	return c.writer.Header()
}

func (c *compressWriter) Write(p []byte) (int, error) {
	written, err := c.zipWriter.Write(p)
	if err == nil {
		c.Header().Set("Content-Length", strconv.Itoa(written))
	}
	return written, err
}

func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.writer.Header().Set("Content-Encoding", "gzip")
	}
	c.writer.WriteHeader(statusCode)
}

// Close закрывает gzip.Writer и досылает все данные из буфера.
func (c *compressWriter) Close() error {
	return c.zipWriter.Close()
}

// compressReader реализует интерфейс io.ReadCloser и позволяет прозрачно для сервера
type compressReader struct {
	readCloser io.ReadCloser
	zipReader  *gzip.Reader
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		readCloser: r,
		zipReader:  zr,
	}, nil
}

func (c compressReader) Read(p []byte) (n int, err error) {
	return c.zipReader.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.readCloser.Close(); err != nil {
		return err
	}
	return c.zipReader.Close()
}

func GzipMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		originalWriter := w

		acceptEncoding := r.Header.Get("Accept-Encoding")
		contentType := r.Header.Get("Content-Type")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		gzipReasonable := contentType == "application/json" || contentType == "text/html"
		if gzipReasonable && supportsGzip {
			compressWriter := newCompressWriter(w)
			originalWriter = compressWriter
			defer compressWriter.Close()
		}

		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if gzipReasonable && sendsGzip {
			compressReader, err := newCompressReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.Body = compressReader
			defer compressReader.Close()
		}

		h.ServeHTTP(originalWriter, r)
	}
}
