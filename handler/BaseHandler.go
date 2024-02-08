package handler

import (
	"log"

	"github.com/ZertyCraft/GoLogger/formater"
	"github.com/ZertyCraft/GoLogger/levels"
)

// Handler is an interface that defines the behavior of a log handler.
type Handler interface {
	Log(level levels.Level, message string)
}

// BaseHandler is a struct that implements the Handler interface.
type BaseHandler struct {
	Handler  // Embed the Handler interface
	formater formater.Formater

	Level levels.Level
}

// `NewBaseHandler` creates a new instance of BaseHandler.
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		Handler:  nil,
		formater: formater.NewBaseFormater(""),
		Level:    levels.INFO,
	}
}

// `SetLevel` sets the level of the handler.
func (h *BaseHandler) SetLevel(level levels.Level) {
	// Check if the level is valid
	if level < levels.DEBUG || level > levels.CRITICAL {
		log.Fatal("Invalid level")
	}

	h.Level = level
}

// `GetLevel` returns the level of the handler.
func (h *BaseHandler) GetLevel() levels.Level {
	return h.Level
}

// `SetFormater` sets the formater of the handler.
func (h *BaseHandler) SetFormater(formater formater.Formater) {
	h.formater = formater
}

// `isLevelSufficient` checks if the given level is sufficient to be logged.
// isLevelSufficient checks if the given log level is sufficient based on the handler's level.
// It returns true if the log level is greater than or equal to the handler's level, otherwise false.
func (h *BaseHandler) isLevelSufficient(level levels.Level) bool {
	return level >= h.Level
}

// `Log` logs the given message using the handler (not implemented in BaseHandler).
func (h *BaseHandler) Log(_ levels.Level, _ string) {
	log.Fatal("`Log` method not implemented in `BaseHandler`")
}
