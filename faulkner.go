// Faulkner is currently a simple wrapper over Go's log package.

package faulkner

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

type Logger struct {
	Debug  *log.Logger
	Info   *log.Logger
	Error  *log.Logger
	Banner *log.Logger
}

type LogOptions struct {
	OutWriter io.Writer
	Debug     bool
	Info      bool
	flags     int
}

// DefaultLogger just wraps NewLogger and takes care of the error value
func DefaultLogger() Logger {
	l, _ := NewLogger() // this can't return an error anyways
	return l
}

// NewLogger creates a new Log instance
func NewLogger(options ...func(*LogOptions) error) (Logger, error) {

	// start by setting sane defaults
	log_options := LogOptions{
		OutWriter: os.Stderr,             // write logging to standard error
		Debug:     true,                  // write debug messages
		Info:      true,                  // write info messages
		flags:     log.Ldate | log.Ltime, // start lines with date and time
	}

	// now loop over options and set any values needed
	for _, option := range options {
		err := option(&log_options)
		if err != nil {
			return Logger{}, err
		}
	}

	// now that we have all the options set, we can create our loggers
	// based on defaults + any options set
	logger := Logger{}

	logger.Banner = log.New(log_options.OutWriter, "", 0)
	if log_options.OutWriter == os.Stderr && runtime.GOOS != "windows" {
		logger.Error = log.New(log_options.OutWriter, "\033[1;31mERROR: \033[0m", log_options.flags|log.Lshortfile)
	} else {
		logger.Error = log.New(log_options.OutWriter, "ERROR: ", log_options.flags|log.Lshortfile)
	}
	if log_options.Info {
		logger.Info = log.New(log_options.OutWriter, "INFO:  ", log_options.flags)
	} else {
		log.New(ioutil.Discard, "", 0)
	}
	if log_options.Debug {
		logger.Debug = log.New(log_options.OutWriter, "DEBUG: ", log_options.flags|log.Lshortfile)
	} else {
		logger.Debug = log.New(ioutil.Discard, "", 0)
	}

	return logger, nil
}

func SetDebug(d bool) func(s *LogOptions) error {
	return func(s *LogOptions) error {
		s.Debug = d
		return nil
	}
}

func SetInfo(d bool) func(s *LogOptions) error {
	return func(s *LogOptions) error {
		s.Info = d
		return nil
	}
}

func SetFile(fn string) func(s *LogOptions) error {
	return func(s *LogOptions) error {
		file, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		s.OutWriter = file
		return nil
	}
}

func SetBuffer(buf *bytes.Buffer) func(s *LogOptions) error {
	return func(s *LogOptions) error {
		s.OutWriter = buf
		return nil
	}
}

func (l *Logger) PrintBanner(message string) {
	l.Banner.Println("--------------------------")
	l.Banner.Println(message)
	l.Banner.Println("--------------------------")
}

func (l *Logger) DebugOff() {
	l.Debug = log.New(ioutil.Discard, "", 0)
}

func (l *Logger) InfoOff() {
	l.Info = log.New(ioutil.Discard, "", 0)
}
