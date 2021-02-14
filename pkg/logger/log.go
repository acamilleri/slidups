package logger

import (
	"github.com/sirupsen/logrus"
)

// Logger - Structure to wrap Logrus
type Logger struct {
	*logrus.Logger
}

// New - Initialize a wrap of Logrus as logger
func New() *Logger {
	return &Logger{Logger: logrus.New()}
}
