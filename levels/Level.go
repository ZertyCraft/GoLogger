package levels

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	CRITICAL
)

// `String` is a method that returns the string representation of the log level.
func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case CRITICAL:
		return "CRITICAL"
	}

	return "UNKNOWN"
}
