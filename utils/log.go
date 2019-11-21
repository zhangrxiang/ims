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

func LoadLogInit() {
	Mkdir("./logs")
	_logger = logrus.New()
	var err error
	_logf, err = rotatelogs.New(
		"./logs/log_%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(30*time.Hour*24),
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

func Debug(args ...interface{}) {
	if _logger.Level >= logrus.DebugLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)

		msg := fmt.Sprint(args...)
		entry.Debug(msg)
	}
}

func Info(args ...interface{}) {
	if _logger.Level >= logrus.InfoLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Info(msg)
	}
}

func Warn(args ...interface{}) {
	if _logger.Level >= logrus.WarnLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Warn(msg)
	}
}

func Error(args ...interface{}) {
	if _logger.Level >= logrus.ErrorLevel {
		entry := _logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		msg := fmt.Sprint(args...)
		entry.Error(msg)
	}
}

func Fatal(args ...interface{}) {
	if _logger.Level >= logrus.FatalLevel {
		entry := _logger.WithFields(logrus.Fields{})
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
