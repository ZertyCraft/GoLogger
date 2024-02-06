package logger

import (
	"github.com/ZertyCraft/GoLogger/handler"
	"github.com/ZertyCraft/GoLogger/levels"
)

type Logger struct {
	handler []handler.Handler
}

func NewLogger() *Logger {
	return &Logger{
		handler: make([]handler.Handler, 0),
	}
}

func (l *Logger) AddHandler(handler handler.Handler) {
	l.handler = append(l.handler, handler)
}

func (l *Logger) RemoveHandler(handler handler.Handler) {
	for i, h := range l.handler {
		if h == handler {
			l.handler = append(l.handler[:i], l.handler[i+1:]...)
		}
	}
}

func (l *Logger) Log(level levels.Level, message string) {
	for _, h := range l.handler {
		h.Log(level, message)
	}
}

func (l *Logger) Debug(message string) {
	l.Log(levels.DEBUG, message)
}

func (l *Logger) Info(message string) {
	l.Log(levels.INFO, message)
}

func (l *Logger) Warning(message string) {
	l.Log(levels.WARN, message)
}

func (l *Logger) Error(message string) {
	l.Log(levels.ERROR, message)
}

func (l *Logger) Critical(message string) {
	l.Log(levels.CRITICAL, message)
}
