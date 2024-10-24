package client

import (
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"

	"github.com/ultradns/ultradns-go-sdk/internal/version"
)

type logLevelType int

const (
	LogOff logLevelType = iota
	LogError
	LogDebug
	LogTrace
)

type logger struct {
	logLevel logLevelType
	logger   *log.Logger
}

func (l logger) getLogPrefix(logLevel logLevelType) string {
	switch logLevel {
	case LogError:
		return "[ERROR] "
	case LogDebug:
		return "[DEBUG] "
	case LogTrace:
		return "[TRACE] "
	}

	return ""
}

func (c *Client) logHttpRequest(req *http.Request) {
	if c.logger.logLevel >= LogDebug && req != nil {
		data, _ := httputil.DumpRequest(req, true)
		c.Debug("HTTP Request Sent:\n\t%s", formatHttpLog(data))
	}
}

func (c *Client) logHttpResponse(res *http.Response) {
	if c.logger.logLevel >= LogDebug && res != nil {
		data, _ := httputil.DumpResponse(res, true)
		c.Debug("HTTP Response Received:\n\t%s", formatHttpLog(data))
	}
}

func formatHttpLog(data []byte) string {
	str := strings.TrimRight(string(data), "\n")
	str = regexp.MustCompile(`\r?\n`).ReplaceAllString(str, "\r\n\t")
	return str
}

func (l logger) logf(logType logLevelType, format string, v ...any) {
	if l.logLevel >= logType && l.logger != nil {
		l.logger.SetPrefix(l.getLogPrefix(logType))
		l.logger.Printf("["+version.GetSDKVersion()+"] "+format+"\n", v...)
	}
}

func (c *Client) Error(format string, v ...any) {
	c.logger.logf(LogError, format, v...)
}

func (c *Client) Debug(format string, v ...any) {
	c.logger.logf(LogDebug, format, v...)
}

func (c *Client) Trace(format string, v ...any) {
	c.logger.logf(LogTrace, format, v...)
}

func (c *Client) EnableDefaultDebugLogger() {
	c.EnableLogger(LogDebug, log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
}

func (c *Client) EnableDefaultTraceLogger() {
	c.EnableLogger(LogTrace, log.Ldate|log.Lmicroseconds|log.Lmsgprefix)
}

func (c *Client) EnableLogger(logLevel logLevelType, flags int) {
	c.logger.logLevel = logLevel
	c.logger.logger = log.Default()
	c.logger.logger.SetFlags(flags)
}

func (c *Client) DisableLogger() {
	c.logger.logLevel = LogOff
}
