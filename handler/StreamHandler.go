package handler

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sync"

	"github.com/ZertyCraft/GoLogger/levels"
)

type StreamHandler struct {
	BaseHandler
	level             levels.Level
	useLock           bool
	filePermission    int
	fileName          string
	filePath          string
	bufferSize        int
	writer            *bufio.Writer
	file              *os.File
	mutex             sync.Mutex
	isFileWriterSetup bool
}

type DefaultConfig struct {
	filePermission    int
	useLock           bool
	level             levels.Level
	logPath           string
	logFileName       string
	bufferSize        int
	isFileWriterSetup bool
}

func NewStreamHandler() *StreamHandler {
	const defaultFilePermission = 0o666

	const bufferSize = 4096

	defaultConfig := DefaultConfig{
		filePermission:    defaultFilePermission,
		useLock:           true,
		level:             levels.INFO,
		logPath:           "logs",
		logFileName:       "log",
		bufferSize:        bufferSize,
		isFileWriterSetup: false,
	}

	handler := &StreamHandler{
		BaseHandler:       *NewBaseHandler(),
		level:             defaultConfig.level,
		useLock:           defaultConfig.useLock,
		filePermission:    defaultConfig.filePermission,
		fileName:          defaultConfig.logFileName,
		filePath:          defaultConfig.logPath,
		bufferSize:        defaultConfig.bufferSize,
		writer:            nil,
		file:              nil,
		mutex:             sync.Mutex{},
		isFileWriterSetup: defaultConfig.isFileWriterSetup,
	}

	return handler
}

func (h *StreamHandler) SetLevel(level levels.Level) {
	h.level = level
}

func (h *StreamHandler) SetFilePermission(permission int) {
	h.filePermission = permission
}

func (h *StreamHandler) SetFileName(name string) {
	h.fileName = name
}

func (h *StreamHandler) SetFilePath(path string) {
	h.filePath = path
}

func (h *StreamHandler) SetBufferSize(size int) {
	h.bufferSize = size
}

func (h *StreamHandler) SetUseLock(useLock bool) {
	h.useLock = useLock
}

func (h *StreamHandler) setupFileWriter() error {
	if err := h.createDir(); err != nil {
		return err
	}

	file, err := h.createFile()
	if err != nil {
		return err
	}

	h.file = file
	h.writer = bufio.NewWriterSize(file, h.bufferSize)

	return nil
}

const directoryPermission = 0o755

func (h *StreamHandler) createDir() error {
	if _, err := os.Stat(h.filePath); os.IsNotExist(err) {
		if err := os.MkdirAll(h.filePath, directoryPermission); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	return nil
}

func (h *StreamHandler) createFile() (*os.File, error) {
	filePath := fmt.Sprintf("%s/%s.log", h.filePath, h.fileName)

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, fs.FileMode(h.filePermission))
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	return file, nil
}

func (h *StreamHandler) Log(level levels.Level, message string) {
	if !h.isFileWriterSetup {
		if err := h.setupFileWriter(); err != nil {
			log.Fatalf("StreamHandler initialization failed: %v", err)
		}
	}

	if h.useLock {
		h.mutex.Lock()
		defer h.mutex.Unlock()
	}

	if level < h.level {
		return
	}

	formattedMessage, err := h.formater.Format(level, message)
	if err != nil {
		log.Printf("Error formatting message: %v\n", err)

		return
	}

	// Add line break if not present
	if formattedMessage[len(formattedMessage)-1] != '\n' {
		formattedMessage += "\n"
	}

	if _, err := h.writer.Write([]byte(formattedMessage)); err != nil {
		log.Printf("Error writing to log file: %v\n", err)

		return
	}

	if err := h.writer.Flush(); err != nil {
		log.Printf("Error flushing buffer: %v\n", err)
	}
}

func (h *StreamHandler) Close() {
	if h.useLock {
		h.mutex.Lock()
		defer h.mutex.Unlock()
	}

	if err := h.writer.Flush(); err != nil {
		log.Printf("Error flushing buffer on close: %v\n", err)
	}

	if err := h.file.Close(); err != nil {
		log.Printf("Error closing log file: %v\n", err)
	}
}
