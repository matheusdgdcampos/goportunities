package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)
	return &Logger{
		debug:   log.New(writer, "[DEBUG]: ", logger.Flags()),
		info:    log.New(writer, "[INFO]: ", logger.Flags()),
		warning: log.New(writer, "[WARNING]: ", logger.Flags()),
		err:     log.New(writer, "[ERROR]: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(args ...any) {
	l.debug.Println(args...)
}

func (l *Logger) Info(args ...any) {
	l.info.Println(args...)
}

func (l *Logger) Warning(args ...any) {
	l.warning.Println(args...)
}

func (l *Logger) Error(args ...any) {
	l.err.Println(args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.debug.Printf(format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.info.Printf(format, args...)
}

func (l *Logger) Warningf(format string, args ...any) {
	l.warning.Printf(format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.err.Printf(format, args...)
}
