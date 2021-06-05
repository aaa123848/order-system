package logger

import (
	"io"
	"os"
)

var errorFile io.Writer

func InitLogFile() {
	f, err := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	errorFile = f
}

func GetErrorLog(err error) {
	errorFile.Write([]byte(err.Error()))
}
