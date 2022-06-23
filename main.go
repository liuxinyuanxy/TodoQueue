// @title        Golang Service Template
// @version      0.1
// @description  Golang back-end service template, get started with back-end projects quickly
// @BasePath     /api

package main

import (
	"TodoQueue/app"
	_ "TodoQueue/docs"
	"TodoQueue/model"
)

type LogFormatter struct{}

func (slf *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("06/01/02 15:04:05")
	msg := fmt.Sprintf("[%s] %s (%s:%d): %s\n", entry.Level,
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
