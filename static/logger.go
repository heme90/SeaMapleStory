package static

import (
	"flag"
	"github.com/google/logger"
	"log"
)

var serviceLogger *logger.Logger

type ServiceLogger struct {
	logger *logger.Logger
}

func InitLogger() {
	ver := flag.Bool("MRNG",false ,"use strict, only Error for Error")
	//TODO if need, add log file by OS.OpenFIle
	logger.SetFlags(log.LstdFlags)
	serviceLogger = logger.Init("MRNG LOGGER",*ver,true,nil)
}

func NewLogger() *ServiceLogger {
	return &ServiceLogger{logger: serviceLogger}
}

//common logging
func (l *ServiceLogger) Info(function, msg string) {
	l.logger.Info("FUNC %s : %s",function, msg)
}

//logging For None Error case Warn
func (l *ServiceLogger) Warn(function, msg string) {
	l.logger.Warning("FUNC %s : %s",function, msg)
}
//logging When something critical
func (l *ServiceLogger) Error(function, msg string) {
	l.logger.Error("FUNC %s : %s",function, msg)
}
//logging for panic, ... or thing may happens service down
func (l *ServiceLogger) Fatal(function, msg string) {
	l.logger.Fatal("FUNC %s : %s",function, msg)
}