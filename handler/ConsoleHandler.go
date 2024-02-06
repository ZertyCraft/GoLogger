package handler

import (
	"io"
	"log"
	"os"

	"github.com/ZertyCraft/GoLogger/levels"
)

// ConsoleHandler is a struct that represents a console log handler.
type ConsoleHandler struct {
	BaseHandler
	logger *log.Logger
}

// NewConsoleHandler creates a new instance of ConsoleHandler.
// If no writer is provided, the default writer is os.Stderr.
func NewConsoleHandler(writer ...io.Writer) *ConsoleHandler {
	var outputWriter io.Writer = os.Stderr

	if len(writer) > 0 {
		outputWriter = writer[0]
	}

	return &ConsoleHandler{
		BaseHandler: *NewBaseHandler(),
		logger:      log.New(outputWriter, "", 0),
	}
}

// Log logs the given message using the console logger.
func (h *ConsoleHandler) Log(level levels.Level, message string) {
	if level >= h.Level {
		formatedMessage, err := h.formater.Format(level, message)
		if err != nil {
			panic(err)
		}

		h.logger.Print(formatedMessage)
	}
}

// SetLoggerOutput sets the output of the logger (default is os.Stdout).
func (h *ConsoleHandler) SetLoggerOutput(output *log.Logger) {
	h.logger = output
}

// GetLoggerOutput returns the output of the logger.
func (h *ConsoleHandler) GetLoggerOutput() *log.Logger {
	return h.logger
}
