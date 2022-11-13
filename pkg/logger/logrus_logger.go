package logger

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

import (
	"encoding/json"
)

type Fields map[string]interface{}

type Entry struct {
	entry *logrus.Entry
}

var loggerLevelMap = map[string]logrus.Level{
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}
func SetLevel(level string) {
	lv, exist := loggerLevelMap[level]
	if !exist {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetLevel(lv)
}

func initFields() Fields {
	return Fields{
		"app": "orc-system",
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

func WithFields(fields Fields) *Entry {
	df := initFields()
	df["data"] = fields
	return &Entry{
		entry: logrus.WithFields(convertToLogrusFields(df)).WithTime(time.Now().UTC()),
	}
}

func WithSecretFields(fields Fields) *Entry {
	df := initFields()
	markData, _ := json.Marshal(fields)
	df["data"] = markData
	return &Entry{
		entry: logrus.WithFields(convertToLogrusFields(df)).WithTime(time.Now().UTC()),
	}
}

func (e *Entry) Debug(args ...interface{}) {
	e.infoLogger().Debug(args...)
}

func (e *Entry) Debugf(template string, args ...interface{}) {
	e.infoLogger().Debugf(template, args...)
}

func (e *Entry) Info(args ...interface{}) {
	e.infoLogger().Info(args...)
}

func (e *Entry) Infof(template string, args ...interface{}) {
	e.infoLogger().Infof(template, args...)
}

func (e *Entry) Printf(template string, args ...interface{}) {
	e.infoLogger().Printf(template, args...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.infoLogger().Warn(args...)
}

func (e *Entry) Warnf(template string, args ...interface{}) {
	e.infoLogger().Warnf(template, args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.infoLogger().Error(args...)
}

func (e *Entry) Errorf(template string, args ...interface{}) {
	e.infoLogger().Errorf(template, args...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.infoLogger().Fatal(args...)
}

func (e *Entry) Fatalf(template string, args ...interface{}) {
	e.infoLogger().Fatalf(template, args...)
}

func (e *Entry) Fatalln(args ...interface{}) {
	e.infoLogger().Fatalln(args...)
}

func (e *Entry) Panic(args ...interface{}) {
	e.infoLogger().Panic(args...)
}

func (e *Entry) Panicf(template string, args ...interface{}) {
	e.infoLogger().Panicf(template, args...)
}

func (e *Entry) Panicln(args ...interface{}) {
	e.infoLogger().Panicln(args...)
}

func defaultWithField() *Entry {
	init := initFields()
	return &Entry{
		entry: logrus.WithFields(convertToLogrusFields(init)).WithTime(time.Now().UTC()),
	}
}

func Debug(args ...interface{}) {
	defaultWithField().infoLogger().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	defaultWithField().infoLogger().Debugf(template, args...)
}

func Info(args ...interface{}) {
	defaultWithField().infoLogger().Info(args...)
}

func Infof(template string, args ...interface{}) {
	defaultWithField().infoLogger().Infof(template, args...)
}

func Printf(template string, args ...interface{}) {
	defaultWithField().infoLogger().Printf(template, args...)
}

func Warn(args ...interface{}) {
	defaultWithField().infoLogger().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	defaultWithField().infoLogger().Warnf(template, args...)
}

func Error(args ...interface{}) {
	defaultWithField().infoLogger().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	defaultWithField().infoLogger().Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	defaultWithField().infoLogger().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	defaultWithField().infoLogger().Fatalf(template, args...)
}

func Fatalln(args ...interface{}) {
	defaultWithField().infoLogger().Fatalln(args...)
}

func Panic(args ...interface{}) {
	defaultWithField().infoLogger().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	defaultWithField().infoLogger().Panicf(template, args...)
}

func Panicln(args ...interface{}) {
	defaultWithField().infoLogger().Panicln(args...)
}

func (e *Entry) infoLogger() *logrus.Entry {
	var loggerFile string
	_, f, _, ok := runtime.Caller(1)
	if ok {
		loggerFile = f
	} else {
		return e.entry
	}
	skip := 2
	for {
		if pc, file, line, ok := runtime.Caller(skip); ok {
			if strings.EqualFold(loggerFile, file) {
				skip++
				continue
			}
			funcName := runtime.FuncForPC(pc).Name()
			slash := strings.LastIndex(file, "/")
			if slash >= 0 {
				file = file[slash+1:]
			}
			return e.entry.WithFields(logrus.Fields{
				"file":     file,
				"function": funcName,
				"line":     line,
			})
		}
		return e.entry
	}
}
