// @title        Golang Service Template
// @version      0.1
// @description  Golang back-end service template, get started with back-end projects quickly
// @BasePath     /api

package main

import (
	"TodoQueue/app"
	_ "TodoQueue/docs"
	"TodoQueue/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

type LogFormatter struct{}

func (slf *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var logColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel, logrus.InfoLevel:
		logColor = 32 //green
	case logrus.ErrorLevel:
		logColor = 31 //red
	case logrus.WarnLevel:
		logColor = 33 //yellow
	default:
		logColor = 36 //blue
	}
	timestamp := time.Now().Local().Format("06/01/02 15:04:05")
	msg := fmt.Sprintf("\x1b[%dm[%s] %s (%s:%d):\x1b[0m %s\n", logColor,
		entry.Level,
		timestamp,
		filepath.Base(entry.Caller.File),
		entry.Caller.Line,
		entry.Message)
	return []byte(msg), nil
}

func main() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})

	model.Init()
	app.InitWebFramework()
	app.StartServer()
}
