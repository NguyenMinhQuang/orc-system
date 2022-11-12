package logger

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

type Fields map[string]interface{}
type Entry struct {
	entry *logrus.Entry
}

func initField() Fields {
	return Fields{
		"app":  "orc-sys",
		"id":   uuid.New().String(),
		"type": "default",
	}
}

func WithFields(fields Fields) *Entry {
	fild := initField()
	return &Entry{
		entry: logrus.WithFields(convertToLogrusFields(fild)).WithTime(time.Now().UTC()),
	}
}

// convertToLogrusFields converts Fields to logrus.Fields
func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}

func (e *Entry) Info(args ...interface{}) {
	e.entry.Info(args...)
}
