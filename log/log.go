package log

import (
	"log"
	"io"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func Init(infoHandle io.Writer, errorHandle io.Writer) {
	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)
	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime)
}
