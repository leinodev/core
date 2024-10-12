package corelogger

import (
	"fmt"
	"log"
)

type CoreLogger interface {
	LogfDebug(format string, args ...interface{})
	LogfError(format string, args ...interface{})
}

type coreLoggerImpl struct {
	CoreLogger
	prefix string
}

func New(parent CoreLogger, prefix string) CoreLogger {
	if parent != nil {
		prefix = fmt.Sprintf("%s %s", parent.(*coreLoggerImpl).prefix, prefix)
	}
	h := &coreLoggerImpl{
		prefix:     prefix,
		CoreLogger: parent,
	}

	return h
}

func (h *coreLoggerImpl) LogfDebug(format string, args ...interface{}) {
	format = fmt.Sprintf("%s %s", h.prefix, format)
	log.Printf(format, args...)
}

func (h *coreLoggerImpl) LogfError(format string, args ...interface{}) {
	format = fmt.Sprintf("%s %s", h.prefix, format)
	log.Printf(format, args...)
}
