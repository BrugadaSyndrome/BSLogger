package bslogger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	logFile   *os.File
	logger    *log.Logger
	Name      string
	Verbosity Verbosity
}

func NewLogger() Logger {
	return Logger{
		logger:    log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmsgprefix),
		Name:      "BSLogger",
		Verbosity: Normal,
	}
}

func (l *Logger) Fatal(message string) {
	if l.logFile != nil {
		l.logger.SetOutput(l.logFile)
		l.logger.SetPrefix(fmt.Sprintf("%s FATAL: %s", l.Name, message))
		l.logger.Print(message)
	}

	l.logger.SetOutput(os.Stderr)
	l.logger.SetPrefix(l.Name + ansiEscapeEncode(" FATAL: ", fgBrightRed, bgDefault, framed))
	l.logger.Fatalf(ansiEscapeEncode(message, fgBrightRed, bgDefault, framed))
}

func (l *Logger) Fatalf(format string, values ...interface{}) {
	l.Fatal(fmt.Sprintf(format, values...))
}

func (l *Logger) Error(message string) {
	if l.Verbosity < Minimal {
		return
	}

	if l.logFile != nil {
		l.logger.SetOutput(l.logFile)
		l.logger.SetPrefix(fmt.Sprintf("%s ERROR: ", l.Name))
		l.logger.Print(message)
	}

	l.logger.SetOutput(os.Stderr)
	l.logger.SetPrefix(l.Name + ansiEscapeEncode(" ERROR: ", fgRed, bgDefault, normal))
	l.logger.Print(ansiEscapeEncode(message, fgRed, bgDefault, normal))
}

func (l *Logger) Errorf(format string, values ...interface{}) {
	l.Error(fmt.Sprintf(format, values...))
}

func (l *Logger) Warning(message string) {
	if l.Verbosity < Normal {
		return
	}

	if l.logFile != nil {
		l.logger.SetOutput(l.logFile)
		l.logger.SetPrefix(fmt.Sprintf("%s WARNING: ", l.Name))
		l.logger.Print(message)
	}

	l.logger.SetOutput(os.Stdout)
	l.logger.SetPrefix(l.Name + ansiEscapeEncode(" WARNING: ", fgYellow, bgDefault, normal))
	l.logger.Print(ansiEscapeEncode(message, fgYellow, bgDefault, normal))
}

func (l *Logger) Warningf(format string, values ...interface{}) {
	l.Warning(fmt.Sprintf(format, values...))
}

func (l *Logger) Info(message string) {
	if l.Verbosity < Normal {
		return
	}

	if l.logFile != nil {
		l.logger.SetOutput(l.logFile)
		l.logger.SetPrefix(fmt.Sprintf("%s INFO: ", l.Name))
		l.logger.Print(message)
	}

	l.logger.SetOutput(os.Stdout)
	l.logger.SetPrefix(l.Name + ansiEscapeEncode(" INFO: ", fgBlue, bgDefault, normal))
	l.logger.Print(ansiEscapeEncode(message, fgBlue, bgDefault, normal))
}

func (l *Logger) Infof(format string, values ...interface{}) {
	l.Info(fmt.Sprintf(format, values...))
}

func (l *Logger) Debug(message string) {
	if l.Verbosity < All {
		return
	}

	if l.logFile != nil {
		l.logger.SetOutput(l.logFile)
		l.logger.SetPrefix(fmt.Sprintf("%s DEBUG: ", l.Name))
		l.logger.Print(message)
	}

	l.logger.SetOutput(os.Stdout)
	l.logger.SetPrefix(l.Name + ansiEscapeEncode(" DEBUG: ", fgPurple, bgDefault, normal))
	l.logger.Print(ansiEscapeEncode(message, fgPurple, bgDefault, normal))
}

func (l *Logger) Debugf(format string, values ...interface{}) {
	l.Debug(fmt.Sprintf(format, values...))
}

func (l *Logger) SetLogFile(filepath string) error {
	logFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	l.logFile = logFile
	return nil
}
