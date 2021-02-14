package logger

import "github.com/sirupsen/logrus"

// SetLevel - Define log level
// Fallback to INFO if level is invalid
func (log *Logger) SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		log.WithError(err).Error("failed to parse log level: fallback to INFO")
		lvl = logrus.InfoLevel
	}

	log.Logger.SetLevel(lvl)
}
