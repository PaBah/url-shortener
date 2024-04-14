package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLoggerWork(t *testing.T) {
	testMessage := `{"test":"message"}`

	type LogRecord struct {
		URI    string `json:"uri"`
		Method string `json:"method"`
		Status int    `json:"status"`
		Size   int    `json:"size"`
	}

	handler := LoggerMiddleware(mock.NewHandlerMock(testMessage, http.StatusOK))

	expectedLog := LogRecord{
		URI:    "/",
		Method: "GET",
		Status: 200,
		Size:   18,
	}
	_ = Initialize("debug")
	var buffer bytes.Buffer
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	writer := bufio.NewWriter(&buffer)

	zap.ReplaceGlobals(zap.New(zapcore.NewCore(encoder, zapcore.AddSync(writer), zapcore.DebugLevel)))

	srv := httptest.NewServer(handler)

	defer srv.Close()

	buf := bytes.NewBuffer(nil)

	r := httptest.NewRequest("GET", srv.URL, buf)
	r.RequestURI = ""
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	require.NoError(t, err)
	require.JSONEq(t, testMessage, string(b))
	writer.Flush()

	var logRecord *LogRecord
	logData := buffer.String()
	logData = logData[strings.Index(logData, "{"):]

	_ = json.NewDecoder(bytes.NewReader([]byte(logData))).Decode(&logRecord)
	require.Equal(t, expectedLog, *logRecord)
}
