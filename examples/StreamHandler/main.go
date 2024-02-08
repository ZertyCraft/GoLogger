package main

import (
	"github.com/ZertyCraft/GoLogger/formater"
	"github.com/ZertyCraft/GoLogger/handler"
	"github.com/ZertyCraft/GoLogger/levels"
	"github.com/ZertyCraft/GoLogger/logger"
)

func main() {
	// Create a formater
	formater := formater.NewLineFormater() // Create a new line formater

	// Set the format of the formater
	formater.SetFormat("%d - %l - %m") // %d: date, %l: level, %m: message

	// Create the handler
	handler := handler.NewStreamHandler() // Create a new stream handler

	// Set the formater of the handler
	handler.SetFormater(formater)

	// Set the level of the handler
	handler.SetLevel(levels.INFO)

	// Configure the stream handler
	const bufferSize = 4096 // 4KB

	const filePermission = 0o644 // rw-r--r--

	handler.SetLogDirectory("logs")           // The log directory will be created if not exists
	handler.SetFileName("stream.log")         // The file will be named "log.log" (extension is added automatically)
	handler.SetBufferSize(bufferSize)         // Set the buffer size to 4096 bytes
	handler.SetFilePermission(filePermission) // Set the file permission to 0644
	handler.SetUseLock(true)                  // Use a lock to write to the file

	// Defer stream handler flush
	defer handler.Flush() // Flush the buffer before the program ends (to write logs that are still in the buffer)

	// Create the logger
	logger := logger.NewLogger()

	// Add the handler to the logger
	logger.AddHandler(handler)

	// Log some messages
	logger.Debug("This is a debug message") // Not logged (level is not sufficient)
	logger.Info("This is an info message")  // Logged (level is sufficient)
	// Output : `2006-01-01 00:00:00 - INFO - This is an info message`

	// Change the format of the formater
	formater.SetFormat("[%l] - %m") // %l: level, %m: message

	// Log some messages
	logger.Warning("This is a warning message") // Output: `[WARN] - This is a warning message`
}
