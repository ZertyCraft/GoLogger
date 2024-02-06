package handler_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/ZertyCraft/GoLogger/formater"
	"github.com/ZertyCraft/GoLogger/handler"
	"github.com/ZertyCraft/GoLogger/levels"
)

// TestConsoleHandler_Log_Logged test the Log method with a level that should be logged.
func TestConsoleHandler_Log_Logged(t *testing.T) {
	t.Parallel()

	// Create a buffer to capture the output of the logger.
	var buf bytes.Buffer

	consoleHandler := handler.NewConsoleHandler(&buf)
	lineFormater := formater.NewLineFormater()

	lineFormater.SetFormat("%l %m")
	consoleHandler.SetFormater(lineFormater)
	consoleHandler.SetLevel(levels.INFO)

	consoleHandler.Log(levels.INFO, "TestConsoleHandler_Log_Logged")

	if got := strings.TrimSpace(buf.String()); got != "INFO TestConsoleHandler_Log_Logged" {
		t.Errorf("Log() = `%v`, want `INFO TestConsoleHandler_Log_Logged`", got)
	}
}

// TestConsoleHandler_Log_NotLogged test the Log method with a level that should not be logged.
func TestConsoleHandler_Log_NotLogged(t *testing.T) {
	t.Parallel()

	// Create a buffer to capture the output of the logger.
	var buf bytes.Buffer

	consoleHandler := handler.NewConsoleHandler(&buf)
	lineFormater := formater.NewLineFormater()

	lineFormater.SetFormat("%l %m")
	consoleHandler.SetFormater(lineFormater)
	consoleHandler.SetLevel(levels.INFO)

	consoleHandler.Log(levels.DEBUG, "TestConsoleHandler_Log_NotLogged")

	if got := strings.TrimSpace(buf.String()); got != "" {
		t.Errorf("Log() = `%v`, want ``", got)
	}
}
