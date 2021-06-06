package logger

import (
	"fmt"
	myerror "order/errors"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogFile() {
	f, err := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	l.Out = f
	Log = l
}

func GetErrorLog(err error) {
	Log.Error(err)
}

func LogStdError(e error) {
	err := myerror.GetStdError(e)
	Log.WithFields(logrus.Fields{
		"funcname":    err.FuncName,
		"happenPoint": fmt.Sprintf("file: %v ; line: %v", err.File, err.Line),
		"callerInfo":  fmt.Sprintf("file: %v ; line: %v", err.CallerFile, err.CallerFuncLine),
	}).Error(err)
}
