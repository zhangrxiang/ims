package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

var (
	logger          = logrus.New()
	TimestampFormat = "2006-01-02 15:04:05"
)

func init() {
	logger.Formatter = &logrus.TextFormatter{
		DisableTimestamp: false,
		FullTimestamp:    true,
		ForceQuote:       true,
		TimestampFormat:  TimestampFormat,
		ForceColors:      true,
	}
}

func Debug(args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Debug(msg)
	}
}
func Info(args ...interface{}) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Info(msg)
	}
}

func Warn(args ...interface{}) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Warn(msg)
	}
}

func Error(args ...interface{}) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Error(msg)
	}
}

func Fatal(args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Fatal(msg)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
