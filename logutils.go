package main

import (
	"log"
	"io"
)

var (
    Debug   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
)

func SetUpLogger(
	debugHandler io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer) {

	Debug = log.New(debugHandler,
        "DEBUG: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Warning = log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}