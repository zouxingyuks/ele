package tools

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var logger *logrus.Logger

func initLog() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logPath := Configs.GetString("logs.path")
	logLevel := Configs.GetString("logs.level")
	logDir := filepath.Dir(logPath)
	// 如果日志文件夹不存在，则创建它
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err = os.MkdirAll(logDir, 0755); err != nil {
			logger.Fatalln(err)
		}
	}

	// 如果日志文件不存在，则创建它
	_, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatalln(err)
	}

	writer, err := rotatelogs.New(
		logPath+".%Y%m%d%H%M%S",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		logger.Fatalln(err)
	}

	logger.SetOutput(writer)

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}
}

func NewLog(mod string) *logrus.Entry {
	return logger.WithField("module", mod)
}
