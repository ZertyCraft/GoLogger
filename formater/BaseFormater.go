package formater

import (
	"log"

	"github.com/ZertyCraft/GoLogger/levels"
)

// Formater is an interface that defines the behavior of a log formatter.
type Formater interface {
	Format(level levels.Level, message string) (string, error)
}

// BaseFormater is a struct that implements the Formater interface.
type BaseFormater struct {
	format string
}

// NewBaseFormater is a function that returns a new BaseFormater.
func NewBaseFormater(format string) *BaseFormater {
	return &BaseFormater{format: format}
}

// SetFormat is a method that sets the format of the log message.
func (f *BaseFormater) SetFormat(format string) {
	f.format = format
}

// GetFormat is a method that returns the format of the log message.
func (f *BaseFormater) GetFormat() string {
	return f.format
}

// Format is a method that formats the log message based on the provided log level and message.
// It returns the formatted log message and an error, if any.
func (f *BaseFormater) Format(_ levels.Level, _ string) (string, error) {
	log.Fatal("`Format` method not implemented in `BaseFormater`")

	return "", nil
}
