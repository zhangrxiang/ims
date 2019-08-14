package utils

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var _logger *logrus.Logger
var _logf *rotatelogs.RotateLogs
var _maxAge int

func InitLog(maxAge int) {
	_maxAge = maxAge
	_logger = logrus.New()
	var err error
	_logf, err = rotatelogs.New(
		"./logs/log_%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour*24),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	_logger.SetFormatter(&logrus.TextFormatter{})
	_logger.SetOutput(os.Stdout)
	_logger.SetLevel(logrus.DebugLevel)
	pathMap := lfshook.WriterMap{
		logrus.InfoLevel:  _logf,
		logrus.WarnLevel:  _logf,
		logrus.ErrorLevel: _logf,
		logrus.FatalLevel: _logf,
		logrus.PanicLevel: _logf,
	}

	lfHook := lfshook.NewHook(pathMap, &logrus.JSONFormatter{})

	_logger.AddHook(lfHook)
}

func ConfigLog(maxAge int) {
	if _maxAge == maxAge {
		return
	}
	if _logger == nil {
		return
	}
	if _logf != nil {
		_ = _logf.Close()
	}
	_logger.ReplaceHooks(map[logrus.Level][]logrus.Hook{})
	logf, err := rotatelogs.New(
		"./logs/log_%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour*24),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	pathMap := lfshook.WriterMap{
		logrus.InfoLevel:  logf,
		logrus.WarnLevel:  logf,
		logrus.ErrorLevel: logf,
		logrus.FatalLevel: logf,
		logrus.PanicLevel: logf,
	}

	lfHook := lfshook.NewHook(pathMap, &logrus.JSONFormatter{})

	_logger.AddHook(lfHook)
	_maxAge = maxAge
}

// Debug
func Debug(args ...interface{}) {
	if _logger.Level >= logrus.DebugLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Debug(msg)
	}
}

/*
// 带有field的Debug
func DebugWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.DebugLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Debug(l)
	}
}*/
// Info
func Info(args ...interface{}) {
	if _logger.Level >= logrus.InfoLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Info(msg)
	}
}

/*
// 带有field的Info
func InfoWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.InfoLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Info(l)
	}
}*/
// Warn
func Warn(args ...interface{}) {
	if _logger.Level >= logrus.WarnLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Warn(msg)
	}
}

/*
// 带有Field的Warn
func WarnWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.WarnLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Warn(l)
	}
}*/
// Error
func Error(args ...interface{}) {
	if _logger.Level >= logrus.ErrorLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Error(msg)
	}
}

/*
// 带有Fields的Error
func ErrorWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.ErrorLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Error(l)
	}
}*/
// Fatal
func Fatal(args ...interface{}) {
	if _logger.Level >= logrus.FatalLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Fatal(msg)
	}
}

/*
// 带有Field的Fatal
func FatalWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.FatalLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Fatal(l)
	}
}*/
/*
// Panic
func Panic(args ...interface{}) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Panic(args...)
	}
}*/
/*
// 带有Field的Panic
func PanicWithFields(l interface{}, f Fields) {
	if logger.Level >= logrus.PanicLevel {
		entry := logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Panic(l)
	}
}*/
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
