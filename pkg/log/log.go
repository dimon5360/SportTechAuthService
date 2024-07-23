package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

type LogLevel = uint8

const (
	none_level LogLevel = iota
	error_level
	warning_level
	info_level
	debug_level
	trace_level
)

const (
	LOG_LEVEL       LogLevel = trace_level
	USE_FILE_OUTPUT bool     = false

	LOGS_FOLDER_KEY          string = "LOGS_FOLDER"
	LOGS_FORMAT_DATETIME_KEY string = "LOGS_FORMAT_DATETIME"
	LOGS_FILENAME_FORMAT     string = "./%s/%s.log"
)

type Interface interface {
	Error(fmt string, v ...any)
	Warning(fmt string, v ...any)
	Info(fmt string, v ...any)
	Debug(fmt string, v ...any)
	Trace(fmt string, v ...any)
}

type logger struct {
	IsInitialized bool
}

var instance logger

func Instance() Interface {

	if instance.IsInitialized {
		return &instance
	}

	if USE_FILE_OUTPUT {

		filename := fmt.Sprintf(LOGS_FILENAME_FORMAT, os.Getenv(LOGS_FOLDER_KEY),
			time.Now().Format(os.Getenv(LOGS_FORMAT_DATETIME_KEY)))

		logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return nil
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		go func() {
			for sig := range c {
				log.Println(sig.String())
				logFile.Close() // close output file after application termination
				os.Exit(0)
			}
		}()

		mw := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(mw)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	instance.Trace("Logger initialized")
	instance.IsInitialized = true

	return &instance
}

func (i *logger) display(fmt string, v ...any) {
	log.Printf(fmt+"\n", v...)
}

func (i *logger) fatal(fmt string, v ...any) {
	log.Fatalf(fmt+"\n", v...)
}

func (i *logger) Error(fmt string, v ...any) {
	if LOG_LEVEL < error_level {
		return
	}
	i.fatal(fmt, v...)
}

func (i *logger) Warning(fmt string, v ...any) {
	if LOG_LEVEL < warning_level {
		return
	}
	i.display(fmt, v...)
}

func (i *logger) Info(fmt string, v ...any) {
	if LOG_LEVEL < info_level {
		return
	}
	i.display(fmt, v...)
}

func (i *logger) Debug(fmt string, v ...any) {
	if LOG_LEVEL < debug_level {
		return
	}
	i.display(fmt, v...)
}

func (i *logger) Trace(fmt string, v ...any) {
	if LOG_LEVEL < trace_level {
		return
	}
	i.display(fmt, v...)
}
