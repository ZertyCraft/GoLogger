package handler_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/ZertyCraft/GoLogger/formater"
	"github.com/ZertyCraft/GoLogger/handler"
	"github.com/ZertyCraft/GoLogger/levels"
)

const (
	logDirectory   = "logs"
	bufferSize     = 32    // 32 bytes
	filePermission = 0o644 // rw-r--r--
	maxFileSize    = 1024  // 1KB
	maxBackupCount = 2     // 2 backup files
)

func initializeRotatingFileHandler(level levels.Level, fileName string) *handler.RotatingFileHandler {
	formater := formater.NewLineFormater()
	formater.SetFormat("%m")

	rotatingFileHandler := handler.NewRotatingFileHandler()
	rotatingFileHandler.SetFormater(formater)
	rotatingFileHandler.SetLevel(level)
	rotatingFileHandler.SetLogDirectory(logDirectory)
	rotatingFileHandler.SetFileName(fileName)
	rotatingFileHandler.SetMaxFileSize(maxFileSize)
	rotatingFileHandler.SetMaxBackupCount(maxBackupCount)
	rotatingFileHandler.SetBufferSize(bufferSize)
	rotatingFileHandler.SetFilePermission(filePermission)
	rotatingFileHandler.SetUseLock(true)

	return rotatingFileHandler
}

func cleanup(filePath string) func() {
	return func() {
		os.Remove(filePath)     // Ignore error for simplicity
		os.Remove(logDirectory) // This might leave the directory if other files are present, which is okay for this example
	}
}

func TestRotatingFileHandler_Log_Logged(t *testing.T) {
	t.Parallel()

	logFileName := "test_logged.log"
	handler := initializeRotatingFileHandler(levels.INFO, logFileName)
	logFilePath := fmt.Sprintf("%s/%s", logDirectory, logFileName)
	t.Cleanup(cleanup(logFilePath))

	handler.Log(levels.INFO, "This is an info message")
	handler.Flush()

	file, err := os.Open(logFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, bufferSize)

	line, _, err := reader.ReadLine()
	if err != nil {
		t.Fatal(err)
	}

	if string(line) != "This is an info message" {
		t.Errorf("Expected log message not found. Got `%v`", string(line))
	}
}

func TestRotatingFileHandler_Log_NotLogged(t *testing.T) {
	t.Parallel()

	logFileName := "test_not_logged.log"
	handler := initializeRotatingFileHandler(levels.INFO, logFileName)
	logFilePath := fmt.Sprintf("%s/%s", logDirectory, logFileName)
	t.Cleanup(cleanup(logFilePath))

	handler.Log(levels.DEBUG, "This is a debug message") // Not logged due to level
	handler.Flush()

	_, err := os.Stat(logFilePath)
	if !os.IsNotExist(err) {
		t.Error("Log file should not have been created")
	}
}
