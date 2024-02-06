package main

import (
	"github.com/ZertyCraft/GoLogger/formater"
	"github.com/ZertyCraft/GoLogger/handler"
	"github.com/ZertyCraft/GoLogger/levels"
	"github.com/ZertyCraft/GoLogger/logger"
)

func main() {
	// Create two line formater
	lineFormaterConsole := formater.NewLineFormater()
	lineFormaterFile := formater.NewLineFormater()

	// Set two different format for the two line formater
	lineFormaterConsole.SetFormat("%d - %l - %m")
	lineFormaterFile.SetFormat("[%d] [%l] %m")

	// Create the two handlers
	consoleHandler := handler.NewConsoleHandler()
	streamHandler := handler.NewStreamHandler()

	// Set the line formater for the console handler
	consoleHandler.SetFormater(lineFormaterConsole)
	streamHandler.SetFormater(lineFormaterFile)

	// Set two different levels for the two handlers
	consoleHandler.SetLevel(levels.DEBUG)
	streamHandler.SetLevel(levels.ERROR)

	// Change the path and name of the file for the stream handler
	streamHandler.SetFilePath("logs") // The logs directory will be created in current directory if it doesn't exist
	streamHandler.SetFileName("log")  // The file will be named "log.log" (extension is added automatically)

	// Create the logger
	logger := logger.NewLogger()

	// Add the two handlers to the logger
	logger.AddHandler(consoleHandler)
	logger.AddHandler(streamHandler)

	// Log some messages
	logger.Debug("This is a debug message")       // Console
	logger.Info("This is an info message")        // Console
	logger.Warning("This is a warning message")   // Console
	logger.Error("This is an error message")      // Console + File
	logger.Critical("This is a critical message") // Console + File
}
