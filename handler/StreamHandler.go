package handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ZertyCraft/GoLogger/levels"
)

// `StreamHandler` is a struct that implements the Handler interface.
type StreamHandler struct {
	BaseHandler
	useLock        bool
	filePermission int
	fileName       string
	logDirectory   string
	bufferSize     int
	writer         *bufio.Writer
	file           *os.File
	mutex          sync.Mutex
}

const (
	// `defaultBufferSize` is the default buffer size for the `StreamHandler`.
	defaultBufferSize = 4096
	// `defaultFilePermission` is the default file permission for the `StreamHandler`.
	defaultFilePermission = 0o644
	// `defaultFileName` is the default file name for the `StreamHandler`.
	defaultFileName = "log"
	// `defaultlogDirectory` is the default log directory for the `StreamHandler`.
	defaultlogDirectory = "logs"
	// `defaultUseLock` is the default value for the `useLock` field of the `StreamHandler`.
	defaultUseLock = true
)

// `NewStreamHandler` is a function that returns a new `StreamHandler` instance.
// NewStreamHandler creates a new instance of StreamHandler.
// It initializes the StreamHandler struct with default values for its fields.
// Returns a pointer to the newly created StreamHandler.
func NewStreamHandler() *StreamHandler {
	return &StreamHandler{
		BaseHandler:    *NewBaseHandler(),
		useLock:        defaultUseLock,
		filePermission: defaultFilePermission,
		fileName:       defaultFileName,
		logDirectory:   defaultlogDirectory,
		bufferSize:     defaultBufferSize,

		writer: nil,
		file:   nil,
		mutex:  sync.Mutex{},
	}
}

// ======== Setters ========
// `SetUseLock` sets the value of the `useLock` field of the `StreamHandler`.
func (handler *StreamHandler) SetUseLock(useLock bool) {
	handler.useLock = useLock
}

// `SetFilePermission` sets the value of the `filePermission` field of the `StreamHandler`.
func (handler *StreamHandler) SetFilePermission(filePermission int) {
	handler.filePermission = filePermission
}

// `SetFileName` sets the value of the `fileName` field of the `StreamHandler`.
func (handler *StreamHandler) SetFileName(fileName string) {
	handler.fileName = fileName
}

// `SetLogDirectory` sets the value of the `logDirectory` field of the `StreamHandler`.
func (handler *StreamHandler) SetLogDirectory(logDirectory string) {
	handler.logDirectory = logDirectory
}

// `SetBufferSize` sets the value of the `bufferSize` field of the `StreamHandler`.
func (handler *StreamHandler) SetBufferSize(bufferSize int) {
	handler.bufferSize = bufferSize
}

// ======== Getters ========
// `GetUseLock` returns the value of the `useLock` field of the `StreamHandler`.
func (handler *StreamHandler) GetUseLock() bool {
	return handler.useLock
}

// `GetFilePermission` returns the value of the `filePermission` field of the `StreamHandler`.
func (handler *StreamHandler) GetFilePermission() int {
	return handler.filePermission
}

// `GetFileName` returns the value of the `fileName` field of the `StreamHandler`.
func (handler *StreamHandler) GetFileName() string {
	return handler.fileName
}

// `GetLogDirectory` returns the value of the `logDirectory` field of the `StreamHandler`.
func (handler *StreamHandler) GetLogDirectory() string {
	return handler.logDirectory
}

// `GetBufferSize` returns the value of the `bufferSize` field of the `StreamHandler`.
func (handler *StreamHandler) GetBufferSize() int {
	return handler.bufferSize
}

// ======== Methods ========
// `isOpened` checks if the file is opened.
// isOpened checks if the StreamHandler's file is open.
// It returns true if the file is open, and false otherwise.
func (handler *StreamHandler) isOpened() bool {
	return handler.file != nil
}

// `open` opens the file for writing.
// open opens the log file for writing.
// It creates the log directory if it doesn't exist and opens the file in append mode.
// If the file already exists, it appends new log entries to it.
// Returns an error if any operation fails.
func (handler *StreamHandler) open() error {
	// Create the log directory if it doesn't exist
	if _, err := os.Stat(handler.logDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(handler.logDirectory, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	// Open the file
	filePath := handler.logDirectory + "/" + handler.fileName

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(handler.filePermission))
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	// Create a new writer
	handler.writer = bufio.NewWriterSize(file, handler.bufferSize)
	handler.file = file

	return nil
}

// `close` closes the file.
// close closes the StreamHandler by flushing the writer and closing the file.
// If the file is not opened, it returns nil.
// Returns an error if flushing the writer or closing the file fails.
func (handler *StreamHandler) close() error {
	// Check if file is opened
	if !handler.isOpened() {
		return nil
	}

	// Flush the writer
	if err := handler.writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	// Close the file
	if err := handler.file.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	handler.file = nil
	handler.writer = nil

	return nil
}

// `Log` logs the given message using the handler.
// Log writes a log message with the specified level.
// If the file is not opened, it will attempt to open it.
// If opening the file fails, an error will be logged and the function will return.
// If a lock is enabled, it will acquire the lock before writing the log message.
// If the log level is not sufficient, the function will return without writing the message.
// The log message will be formatted using the specified formatter.
// If formatting the message fails, an error will be logged and the function will return.
// If the formatted message does not end with a line break, it will be added.
// The formatted message will be written to the file.
// If writing the message fails, an error will be logged and the function will return.
func (handler *StreamHandler) Log(level levels.Level, message string) {
	if !handler.isOpened() {
		if err := handler.open(); err != nil {
			log.Printf("Failed to open file: %v\n", err)

			return
		}
	}

	// Acquire the lock
	if handler.useLock {
		handler.mutex.Lock()
		defer handler.mutex.Unlock()
	}

	// Check if the level is sufficient
	if !handler.isLevelSufficient(level) {
		return
	}

	// Format the message
	formattedMessage, err := handler.formater.Format(level, message)
	if err != nil {
		log.Printf("Failed to format message: %v\n", err)

		return
	}

	// Add line break if not present
	if formattedMessage[len(formattedMessage)-1] != '\n' {
		formattedMessage += "\n"
	}

	// Write the message
	if _, err := handler.writer.WriteString(formattedMessage); err != nil {
		log.Println("Failed to write message:", err)

		return
	}
}

// `Flush` flushes the writer.
// Flush flushes the writer and ensures that all buffered data is written to the underlying file.
// If the file is not already opened, it will be opened before flushing.
// If the StreamHandler is configured to use a lock, it will acquire the lock before flushing.
// Returns an error if there was a problem flushing the writer or opening the file.
func (handler *StreamHandler) Flush() error {
	// Check if file is opened
	if !handler.isOpened() {
		// Open the file
		if err := handler.open(); err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
	}

	// Acquire the lock
	if handler.useLock {
		handler.mutex.Lock()
		defer handler.mutex.Unlock()
	}

	// Flush the writer
	if err := handler.writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	return nil
}
