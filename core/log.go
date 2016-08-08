package core

import (
	"github.com/go-ozzo/ozzo-log"
)

var Logger = log.NewLogger()

func InitLogging() {
	targetLogFile := log.NewFileTarget()
	targetLogConsole := log.NewConsoleTarget()
	targetLogFile.FileName = "GTams-Server.log"
	targetLogFile.BackupCount = 2
	//targetLogConsole.MaxLevel = log.LevelError
	targetLogConsole.ColorMode = true
	Logger.Targets = append(Logger.Targets, targetLogConsole, targetLogFile)
}