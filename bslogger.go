package bslogger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	logFile   *os.File
	name      string
	verbosity verbosity

	Logger *log.Logger
}

func NewLogger(name string, verbosity verbosity, logFile *os.File) Logger {
	logger := Logger{
		name:      fmt.Sprintf("[%s] ", name),
		verbosity: verbosity,
		Logger:    log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmsgprefix),

		logFile: logFile,
	}
	return logger
}

func (l *Logger) Fatal(message string) {
	if l.logFile != nil {
		l.Logger.SetOutput(l.logFile)
		l.Logger.SetPrefix(fmt.Sprintf("%sFATAL: %s", l.name, message))
		l.Logger.Print(message)
	}

	l.Logger.SetOutput(os.Stderr)
	l.Logger.SetPrefix(l.name + ansiEscapeEncode("FATAL: ", fgBrightRed, bgDefault, framed))
	l.Logger.Fatalf(ansiEscapeEncode(message, fgBrightRed, bgDefault, framed))
}

func (l *Logger) Fatalf(format string, values ...interface{}) {
	l.Fatal(fmt.Sprintf(format, values...))
}

func (l *Logger) Error(message string) {
	if l.verbosity < Minimal {
		return
	}

	if l.logFile != nil {
		l.Logger.SetOutput(l.logFile)
		l.Logger.SetPrefix(fmt.Sprintf("%sERROR: ", l.name))
		l.Logger.Print(message)
	}

	l.Logger.SetOutput(os.Stderr)
	l.Logger.SetPrefix(l.name + ansiEscapeEncode("ERROR: ", fgRed, bgDefault, normal))
	l.Logger.Print(ansiEscapeEncode(message, fgRed, bgDefault, normal))
}

func (l *Logger) Errorf(format string, values ...interface{}) {
	l.Error(fmt.Sprintf(format, values...))
}

func (l *Logger) Warning(message string) {
	if l.verbosity < Normal {
		return
	}

	if l.logFile != nil {
		l.Logger.SetOutput(l.logFile)
		l.Logger.SetPrefix(fmt.Sprintf("%sWARNING: ", l.name))
		l.Logger.Print(message)
	}

	l.Logger.SetOutput(os.Stdout)
	l.Logger.SetPrefix(l.name + ansiEscapeEncode("WARNING: ", fgYellow, bgDefault, normal))
	l.Logger.Print(ansiEscapeEncode(message, fgYellow, bgDefault, normal))
}

func (l *Logger) Warningf(format string, values ...interface{}) {
	l.Warning(fmt.Sprintf(format, values...))
}

func (l *Logger) Info(message string) {
	if l.verbosity < Normal {
		return
	}

	if l.logFile != nil {
		l.Logger.SetOutput(l.logFile)
		l.Logger.SetPrefix(fmt.Sprintf("%sINFO: ", l.name))
		l.Logger.Print(message)
	}

	l.Logger.SetOutput(os.Stdout)
	l.Logger.SetPrefix(l.name + ansiEscapeEncode("INFO: ", fgBlue, bgDefault, normal))
	l.Logger.Print(ansiEscapeEncode(message, fgBlue, bgDefault, normal))
}

func (l *Logger) Infof(format string, values ...interface{}) {
	l.Info(fmt.Sprintf(format, values...))
}

func (l *Logger) Debug(message string) {
	if l.verbosity < All {
		return
	}

	if l.logFile != nil {
		l.Logger.SetOutput(l.logFile)
		l.Logger.SetPrefix(fmt.Sprintf("%sDEBUG: ", l.name))
		l.Logger.Print(message)
	}

	l.Logger.SetOutput(os.Stdout)
	l.Logger.SetPrefix(l.name + ansiEscapeEncode("DEBUG: ", fgPurple, bgDefault, normal))
	l.Logger.Print(ansiEscapeEncode(message, fgPurple, bgDefault, normal))
}

func (l *Logger) Debugf(format string, values ...interface{}) {
	l.Debug(fmt.Sprintf(format, values...))
}
