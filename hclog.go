// Package logrhclog defines an implementation of the github.com/go-logr/logr
// interfaces built on top of github.com/hashicorp/go-hclog
//
// Usage
//
// WIP
package logrhclog

import (
	"github.com/go-logr/logr"
	"github.com/hashicorp/go-hclog"
)

type hclogger struct {
	l       hclog.Logger
	logFunc func(msg string, keysAndValues ...interface{})
}

// implements logr.Logger to be an adapter
var _ logr.Logger = &hclogger{}

// Enabled returns true if any log leve is enabled
func (h *hclogger) Enabled() bool {
	return h.l.IsError() || h.l.IsWarn() || h.l.IsInfo() || h.l.IsDebug() || h.l.IsTrace()
}

// Info prints info message, key must be string, value can be other objects
func (h *hclogger) Info(msg string, keysAndValues ...interface{}) {
	h.logFunc(msg, keysAndValues...)
}

// Error prints error message, key must be string, value can be other objects
func (h *hclogger) Error(msg string, keysAndValues ...interface{}) {
	h.l.Error(msg, keysAndValues...)
}

// V changes logger level, 1~2 to Debug, > 2 for Trace, any other will return to Info
func (h *hclogger) V(level int) logr.Logger {
	switch level {
	case 1, 2:
		return &hclogger{l: logger, logFunc: logger.Debug}
	case > 2:
		return &hclogger{l: logger, logFunc: logger.Trace}
	default:
		// if not defined will return the default
		return NewLogger(h.l)
	}
}

// WithValues returns a new logger with values
func (h *hclogger) WithValues(keysAndValues ...interface{}) logr.Logger {
	return NewLogger(h.l.With(keysAndValues...))
}

// WithName returns a new named logger
func (h *hclogger) WithName(name string) logr.Logger {
	return NewLogger(h.l.Named(name))
}

// NewLogger builds a new logger
func NewLogger(logger hclog.Logger) logr.Logger {
	return &hclogger{l: logger, logFunc: logger.Info}
}
