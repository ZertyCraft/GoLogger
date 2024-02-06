package formater

import (
	"strings"
	"time"

	"github.com/ZertyCraft/GoLogger/levels"
)

type LineFormater struct {
	BaseFormater
}

// NewLineFormater creates a new LineFormater with the default format ("%d %l %m").
func NewLineFormater() *LineFormater {
	return &LineFormater{
		BaseFormater: *NewBaseFormater(
			"%d %l %m",
		),
	}
}

// Format formats the message using the given format (or the default one if not set).
// The format can contain the following placeholders:
// - %d: the current date and time (in the format "2006-01-02 15:04:05").
// - %l: the log level.
// - %m: the message.
func (f *LineFormater) Format(level levels.Level, message string) (string, error) {
	formatedMessage := f.format
	formatedMessage = strings.ReplaceAll(formatedMessage, "%d", time.Now().Format("2006-01-02 15:04:05"))
	formatedMessage = strings.ReplaceAll(formatedMessage, "%l", level.String())
	formatedMessage = strings.ReplaceAll(formatedMessage, "%m", message)

	return formatedMessage, nil
}
