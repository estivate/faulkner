// Faulkner is currently a simple wrapper over Go's log package.

package faulkner

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Logger struct {
	LogDebug  *log.Logger
	LogInfo   *log.Logger
	LogError  *log.Logger
	OutWriter io.Writer
	Debug     bool
	Flags     int
}

// NewLogger creates a new Log instance
func NewLogger(options ...func(*Logger) error) (Logger, error) {

	// start by setting sane defaults
	logger := Logger{
		OutWriter: os.Stderr,             // write logging to standard error
		Debug:     true,                  // write all messages, including debug
		Flags:     log.Ldate | log.Ltime, // start lines with date and time
	}

	// now loop over options and set any values needed
	for _, option := range options {
		err := option(&logger)
		if err != nil {
			return Logger{}, err
		}
	}

	// now that we have all the options set, we can create our loggers
	// based on defaults + any options set
	logger.LogInfo = log.New(logger.OutWriter, "INFO: ", logger.Flags)
	if logger.OutWriter == os.Stderr {
		logger.LogError = log.New(logger.OutWriter, "\033[1;31mERROR: \033[0m", logger.Flags|log.Lshortfile)
	} else {
		logger.LogError = log.New(logger.OutWriter, "ERROR: ", logger.Flags|log.Lshortfile)
	}

	if logger.Debug {
		logger.LogDebug = log.New(logger.OutWriter, "DEBUG: ", logger.Flags|log.Lshortfile)
	} else {
		logger.LogDebug = log.New(ioutil.Discard, "DEBUG: ", logger.Flags)
	}

	return logger, nil
}

func SetDebug(d bool) func(s *Logger) error {
	return func(s *Logger) error {
		s.Debug = d
		return nil
	}
}

func SetFile(fn string) func(s *Logger) error {
	return func(s *Logger) error {
		file, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		s.OutWriter = file
		return nil
	}
}
