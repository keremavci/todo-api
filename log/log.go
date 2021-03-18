package log

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	Logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: false,
	}
}
/*
var Logger *logrus.Logger

type StructuredLogger struct {
	Logger *logrus.Logger
}


func NewStructuredLogger() *StructuredLogger {
	Logger = logrus.New()

	Logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: false,
	}
	sl := &StructuredLogger{Logger: Logger}
	return sl

}

func (log *StructuredLogger) Panicf(msg string,args ...interface{}){
	log.Logger.Panicf(msg, args...)
}

func (log *StructuredLogger) Errorf(msg string,args ...interface{}){
	log.Logger.Errorf(msg, args...)
}

func (log *StructuredLogger) Fatalf(msg string,args ...interface{}){
	log.Logger.Fatalf(msg, args...)
}

func (log *StructuredLogger) Infof(msg string,args ...interface{}){
	log.Logger.Infof(msg, args...)
}

func (log *StructuredLogger) Info(msg string){
	log.Logger.Info(msg)
}
*/
