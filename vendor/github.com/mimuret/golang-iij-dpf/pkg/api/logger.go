package api

import (
	"io"
	"log"
	"os"
)

// logger interface.
type Logger interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

var _ Logger = &StdLogger{}

// simple logger
// change better logger as you like.
type StdLogger struct {
	*log.Logger
	LogLevel int
}

func NewStdLogger(out io.Writer, prefix string, flag int, level int) *StdLogger {
	if out == nil {
		out = os.Stderr
	}
	if prefix == "" {
		prefix = "dpf-client"
	}
	return &StdLogger{
		Logger:   log.New(out, prefix, flag),
		LogLevel: level,
	}
}

func (s *StdLogger) Tracef(format string, args ...interface{}) {
	if s.LogLevel <= 0 {
		s.Printf(format, args...)
	}
}

func (s *StdLogger) Debugf(format string, args ...interface{}) {
	if s.LogLevel <= 1 {
		s.Printf(format, args...)
	}
}

func (s *StdLogger) Infof(format string, args ...interface{}) {
	if s.LogLevel <= 2 {
		s.Printf(format, args...)
	}
}

func (s *StdLogger) Errorf(format string, args ...interface{}) {
	if s.LogLevel <= 4 {
		s.Printf(format, args...)
	}
}
