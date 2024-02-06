package formater

import (
	"log"

	"github.com/ZertyCraft/GoLogger/levels"
)

type Formater interface {
	Format(level levels.Level, message string) (string, error)
}

type BaseFormater struct {
	format string
}

func NewBaseFormater(format string) *BaseFormater {
	return &BaseFormater{format: format}
}

func (f *BaseFormater) SetFormat(format string) {
	f.format = format
}

func (f *BaseFormater) GetFormat() string {
	return f.format
}

func (f *BaseFormater) Format(_ levels.Level, _ string) (string, error) {
	log.Fatal("`Format` method not implemented in `BaseFormater`")

	return "", nil
}
