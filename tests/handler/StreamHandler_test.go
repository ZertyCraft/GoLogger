package handler_test

import (
	"bufio"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/ZertyCraft/GoLogger/formater"
	"github.com/ZertyCraft/GoLogger/handler"
	"github.com/ZertyCraft/GoLogger/levels"
)

// TestStreamHandler_Log_Logged test the Log method with a level that should be logged in file.
func TestStreamHandler_Log_Logged(t *testing.T) {
	t.Parallel()

	streamHandler := handler.NewStreamHandler()
	lineFormater := formater.NewLineFormater()

	lineFormater.SetFormat("%l %m")
	streamHandler.SetFormater(lineFormater)
	streamHandler.SetLevel(levels.INFO)

	// Set the logs directory
	streamHandler.SetFilePath("logs_test_logged")
	streamHandler.SetFileName("test_logged")

	streamHandler.Log(levels.INFO, "TestStreamHandler_Log_Logged")

	// Open the log file
	file, err := os.Open("logs_test_logged/test_logged.log")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Read last log line from file with bufio
	reader := bufio.NewReaderSize(file, 4096)

	line, _, err := reader.ReadLine()
	if err != nil {
		t.Fatal(err)
	}

	if string(line) != "INFO TestStreamHandler_Log_Logged" {
		t.Errorf("Log() = `%v`, want `INFO TestStreamHandler_Log_Logged`", string(line))
	}

	// Remove the log file
	if err := os.Remove("logs_test_logged/test_logged.log"); err != nil {
		t.Fatal(err)
	}

	// Remove the logs directory
	if err := os.Remove("logs_test_logged"); err != nil {
		t.Fatal(err)
	}
}

// TestStreamHandler_Log_NotLogged test the Log method with a level that should not be logged in file.
func TestStreamHandler_Log_NotLogged(t *testing.T) {
	t.Parallel()

	streamHandler := handler.NewStreamHandler()
	lineFormater := formater.NewLineFormater()

	lineFormater.SetFormat("%l %m")
	streamHandler.SetFormater(lineFormater)
	streamHandler.SetLevel(levels.INFO)

	// Set the logs directory
	streamHandler.SetFilePath("logs_test_not_logged")
	streamHandler.SetFileName("test_not_logged")

	streamHandler.Log(levels.DEBUG, "TestStreamHandler_Log_NotLogged")

	// Open the log file
	file, err := os.Open("logs_test_not_logged/test_not_logged.log")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Read last log line from file with bufio
	reader := bufio.NewReaderSize(file, 4096)

	line, _, err := reader.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) {
			line = []byte("") // Set an empty line if end of file is reached
		} else {
			t.Fatal(err)
		}
	}

	if string(line) != "" {
		t.Errorf("Log() = `%v`, want ``", string(line))
	}

	// Remove the log file
	if err := os.Remove("logs_test_not_logged/test_not_logged.log"); err != nil {
		t.Fatal(err)
	}

	// Remove the logs directory
	if err := os.Remove("logs_test_not_logged"); err != nil {
		t.Fatal(err)
	}
}
