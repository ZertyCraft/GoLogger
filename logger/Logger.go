package logger

import (
	"github.com/ZertyCraft/GoLogger/handler"
	"github.com/ZertyCraft/GoLogger/levels"
)

// `Logger` is a struct that contains a slice of handlers.
type Logger struct {
	handler []handler.Handler
}

// `NewLogger` is a function that returns a new instance of Logger.
func NewLogger() *Logger {
	return &Logger{
		handler: make([]handler.Handler, 0),
	}
}

// `AddHandler` is a method that adds a handler to the logger.
func (l *Logger) AddHandler(handler handler.Handler) {
	l.handler = append(l.handler, handler)
}

// `RemoveHandler` is a method that removes a handler from the logger.
func (l *Logger) RemoveHandler(handler handler.Handler) {
	for i, h := range l.handler {
		if h == handler {
			l.handler = append(l.handler[:i], l.handler[i+1:]...)
		}
	}
}

// `Log` is a method that logs a message with the provided log level.
func (l *Logger) Log(level levels.Level, message string) {
	for _, h := range l.handler {
		h.Log(level, message)
	}
}

// `Debug` is a method that logs a message with the DEBUG log level.
func (l *Logger) Debug(message string) {
	l.Log(levels.DEBUG, message)
}

// `Info` is a method that logs a message with the INFO log level.
func (l *Logger) Info(message string) {
	l.Log(levels.INFO, message)
}

// `Warning` is a method that logs a message with the WARN log level.
func (l *Logger) Warning(message string) {
	l.Log(levels.WARN, message)
}

// `Error` is a method that logs a message with the ERROR log level.
func (l *Logger) Error(message string) {
	l.Log(levels.ERROR, message)
}

// `Critical` is a method that logs a message with the CRITICAL log level.
func (l *Logger) Critical(message string) {
	l.Log(levels.CRITICAL, message)
}
