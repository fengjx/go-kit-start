package logger

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fengjx/go-halo/logger"

	"github.com/fengjx/go-kit-start/common/config"
	"github.com/fengjx/go-kit-start/common/env"
)

var Log logger.Logger
var _logDir = filepath.Join("/var", "log")

func init() {
	logConfig := config.GetConfig().Log
	if strings.ToLower(logConfig.Appender) == "console" {
		Log = logger.NewConsole()
		return
	}
	_logDir = logConfig.LogDir
	err := os.MkdirAll(_logDir, 0644)
	if err != nil {
		panic(err)
	}
	initFileLog(logConfig.Level, _logDir, logConfig.MaxSize, logConfig.MaxDays)
}

func initFileLog(level string, logDir string, maxSize, maxDays int) {
	err := os.MkdirAll(logDir, 0644)
	if err != nil {
		panic(err)
	}
	app := env.GetAppName()
	logfile := filepath.Join(logDir, app, app+".log")
	logLevel := logger.GetLogLevel(level)
	Log = logger.New(logLevel, logfile, maxSize, maxDays)
	log.Println("log file", logfile)
	Log.Infof("log file: %s", logfile)
}

func GetLogDir() string {
	return _logDir
}
