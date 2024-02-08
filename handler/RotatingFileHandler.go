package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ZertyCraft/GoLogger/levels"
)

type RotatingFileHandler struct {
	StreamHandler
	maxFileSize    int    // The maximum size of the file before it is rotated in bytes
	maxBackupCount int    // The number of backup files to keep
	filenameFormat string // The format of the filename
}

const (
	// `defaultMaxFileSize` is the default value for the `maxFileSize` field of the `RotatingFileHandler`.
	defaultMaxFileSize = 1000000
	// `defaultMaxBackupCount` is the default value for the `maxBackupCount` field of the `RotatingFileHandler`.
	defaultMaxBackupCount = 5
	// `defaultFilenameFormat` is the default value for the `filenameFormat` field of the `RotatingFileHandler`.
	defaultFilenameFormat = "%s.%d"
)

// `NewRotatingFileHandler` is a function that returns a new `RotatingFileHandler` instance.
func NewRotatingFileHandler() *RotatingFileHandler {
	return &RotatingFileHandler{
		StreamHandler:  *NewStreamHandler(),
		maxFileSize:    defaultMaxFileSize,
		maxBackupCount: defaultMaxBackupCount,
		filenameFormat: defaultFilenameFormat,
	}
}

// ======== Setters ========
// `SetMaxFileSize` sets the value of the `maxFileSize` field of the `RotatingFileHandler`.
func (handler *RotatingFileHandler) SetMaxFileSize(maxFileSize int) {
	handler.maxFileSize = maxFileSize
}

// `SetMaxBackupCount` sets the value of the `maxBackupCount` field of the `RotatingFileHandler`.
func (handler *RotatingFileHandler) SetMaxBackupCount(maxBackupCount int) {
	handler.maxBackupCount = maxBackupCount
}

// `SetFilenameFormat` sets the value of the `filenameFormat` field of the `RotatingFileHandler`.
// Placeholders:
// - %s: the original filename
// - %d: the backup number
// - %t: the current time in the format "2006-01-02T15:04:05"
// - %n: the current time in the format "20060102150405"
// - %d: the current date in the format "2006-01-02"
// - %t: the current time in the format "15:04:05".
func (handler *RotatingFileHandler) SetFilenameFormat(filenameFormat string) {
	handler.filenameFormat = filenameFormat
}

// ======== Getters ========
// `GetMaxFileSize` returns the value of the `maxFileSize` field of the `RotatingFileHandler`.
func (handler *RotatingFileHandler) GetMaxFileSize() int {
	return handler.maxFileSize
}

// `GetMaxBackupCount` returns the value of the `maxBackupCount` field of the `RotatingFileHandler`.
func (handler *RotatingFileHandler) GetMaxBackupCount() int {
	return handler.maxBackupCount
}

// `GetFilenameFormat` returns the value of the `filenameFormat` field of the `RotatingFileHandler`.
func (handler *RotatingFileHandler) GetFilenameFormat() string {
	return handler.filenameFormat
}

// ======== Methods ========

// `getNewFileName` returns the new filename from file name format placeholders.
// Placeholders:
// - %s: the original filename
// - %d: the backup number
// - %t: the current time in the format "2006-01-02T15:04:05"
// - %n: the current time in the format "20060102150405"
// - %d: the current date in the format "2006-01-02"
// - %t: the current time in the format "15:04:05".
func (handler *RotatingFileHandler) getNewFileName(fileName string, backupNumber int) string {
	// Replace the placeholders
	newFileName := handler.filenameFormat
	newFileName = strings.ReplaceAll(newFileName, "%s", fileName)
	newFileName = strings.ReplaceAll(newFileName, "%d", strconv.Itoa(backupNumber))
	newFileName = strings.ReplaceAll(newFileName, "%t", time.Now().Format("2006-01-02T15:04:05"))
	newFileName = strings.ReplaceAll(newFileName, "%n", time.Now().Format("20060102150405"))
	newFileName = strings.ReplaceAll(newFileName, "%d", time.Now().Format("2006-01-02"))
	newFileName = strings.ReplaceAll(newFileName, "%t", time.Now().Format("15:04:05"))

	return newFileName
}

// `renameFile` renames the file.
func (handler *RotatingFileHandler) renameFile(directory string, oldName string, newName string) error {
	err := os.Rename(filepath.Join(directory, oldName), filepath.Join(directory, newName))
	if err != nil {
		return fmt.Errorf("failed to rename file: %w", err)
	}

	return nil
}

// `deleteFile` deletes the file.
func (handler *RotatingFileHandler) deleteFile(directory string, name string) error {
	err := os.Remove(filepath.Join(directory, name))
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// `rotate` rotates the file.
// It renames the file to the new filename
// and creates a new file with the original name.
// Then it deletes the old backup files.
func (handler *RotatingFileHandler) rotate() {
	// Close the file
	handler.close()

	// Get new file name
	newFileName := handler.getNewFileName(handler.fileName, 1)

	// If the file already exists, rename it
	if _, err := os.Stat(filepath.Join(handler.logDirectory, newFileName)); err == nil {
		for i := 1; ; i++ {
			newFileName = handler.getNewFileName(handler.fileName, i)
			if _, err := os.Stat(filepath.Join(handler.logDirectory, newFileName)); err != nil {
				break
			}
		}
	}

	// Rename the file
	if err := handler.renameFile(handler.logDirectory, handler.fileName, newFileName); err != nil {
		panic(err)
	}

	// Create a new file
	if err := handler.open(); err != nil {
		panic(err)
	}
}

// `getFileSize` returns the size of the file in bytes.
func (handler *RotatingFileHandler) getFileSize() int {
	fileInfo, err := os.Stat(filepath.Join(handler.logDirectory, handler.fileName))
	if err != nil {
		panic(err)
	}

	return int(fileInfo.Size())
}

// `sortBackupFiles` sorts the backup files by the backup number.
func (handler *RotatingFileHandler) sortBackupFiles(backupFiles []string) {
	sort.Slice(backupFiles, func(i, j int) bool {
		number1, err1 := strconv.Atoi(strings.Split(backupFiles[i], ".")[len(strings.Split(backupFiles[i], "."))-1])
		number2, err2 := strconv.Atoi(strings.Split(backupFiles[j], ".")[len(strings.Split(backupFiles[j], "."))-1])

		if err1 != nil || err2 != nil {
			log.Fatal("Error parsing backup file numbers:", err1, err2)
		}

		return number1 < number2
	})
}

// `Log` logs the given message using the handler.
func (handler *RotatingFileHandler) Log(level levels.Level, message string) {
	if !handler.isLevelSufficient(level) {
		return
	}

	handler.ensureFileOpened()

	handler.checkAndRotateFile()

	handler.cleanupOldBackups()

	// Log the message using the stream handler
	handler.StreamHandler.Log(level, message)
}

// Checks the file size and rotates the log file if necessary.
func (handler *RotatingFileHandler) checkAndRotateFile() {
	if handler.getFileSize() > handler.maxFileSize {
		log.Println("Rotating file (size =", handler.getFileSize(), ")")
		handler.rotate()
	}
}

// Cleans up old backup files exceeding the maximum backup count.
func (handler *RotatingFileHandler) cleanupOldBackups() {
	files, err := os.ReadDir(handler.logDirectory)
	if err != nil {
		log.Fatal("Error reading log directory:", err)
	}

	backupFiles := make([]string, 0)

	for _, file := range files {
		if strings.HasPrefix(file.Name(), handler.fileName+".") {
			backupFiles = append(backupFiles, file.Name())
		}
	}

	if len(backupFiles) > handler.maxBackupCount {
		handler.sortBackupFiles(backupFiles)

		for i := 0; i < len(backupFiles)-handler.maxBackupCount; i++ {
			if err := handler.deleteFile(handler.logDirectory, backupFiles[i]); err != nil {
				log.Fatal("Error deleting old backup file:", err)
			}
		}
	}
}
