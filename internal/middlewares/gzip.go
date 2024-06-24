package middlewares

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

// Header - returns http.Header
func (c *compressWriter) Header() http.Header {
	return c.writer.Header()
}

// Write - write HTTP response body
func (c *compressWriter) Write(p []byte) (int, error) {
	written, err := c.zipWriter.Write(p)
	if err == nil {
		c.Header().Set("Content-Length", strconv.Itoa(written))
	}
	return written, err
}

// WriteHeader - write HTTP response Header
func (c *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		c.writer.Header().Set("Content-Encoding", "gzip")
	}
	c.writer.WriteHeader(statusCode)
}

// Close - close gzip.Writer and send rest from buffer
func (c *compressWriter) Close() error {
	return c.zipWriter.Close()
}

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

// Read - read compressed request body
func (c compressReader) Read(p []byte) (n int, err error) {
	return c.zipReader.Read(p)
}

// Close - close compress reader
func (c *compressReader) Close() error {
	if err := c.readCloser.Close(); err != nil {
		return err
	}
	return c.zipReader.Close()
}

// GzipMiddleware - middleware for compress request/responses processing
func GzipMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		if sendsGzip {
			compressReader, err := newCompressReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.Body = compressReader
			defer compressReader.Close()
		}

		h.ServeHTTP(originalWriter, r)
	})
}
