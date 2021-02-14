package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// SetFormatter - Define log format
// Fallback to TEXT if format is invalid
func (log *Logger) SetFormatter(format string) {
	switch {
	case strings.EqualFold("TEXT", format):
		log.Formatter = &logrus.TextFormatter{}
		return

	case strings.EqualFold("JSON", format):
		log.Formatter = &logrus.JSONFormatter{}
		return

	default:
		log.Warn("failed to parse log format: fallback to TEXT")
		log.Formatter = &logrus.TextFormatter{}
	}

}
