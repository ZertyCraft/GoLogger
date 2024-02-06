# GoLogger

A simple and versatile logging library for Go. This library provides a customizable way of logging messages with various options such as setting the log level, formatting the output, and specifying multiple outputs.

## Features

* Customizable log formats
* Multiple log levels
* Support for multiple outputs
* Automatic creation of log directories

## Getting Started

To get started using this library, first import it into your project:

```go
import (
    "github.com/ZertyCraft/GoLogger/formater"
    "github.com/ZertyCraft/GoLogger/handler"
    "github.com/ZertyCraft/GoLogger/levels"
    "github.com/ZertyCraft/GoLogger/logger"
)
```

Next, create a new `LineFormater` object and set its format:

```go
lineFormaterConsole := formater.NewLineFormater()
lineFormaterConsole.SetFormat("%d - %l - %m")

lineFormaterFile := formater.NewLineFormater()
lineFormaterFile.SetFormat("[%d] [%l] %m")
```

Then, create one or more `Handler` objects and specify their log level, formatter, and output destination:

```go
consoleHandler := handler.NewConsoleHandler()
consoleHandler.SetFormater(lineFormaterConsole)
consoleHandler.SetLevel(levels.DEBUG)

streamHandler := handler.NewStreamHandler()
streamHandler.SetFormater(lineFormaterFile)
streamHandler.SetLevel(levels.ERROR)
streamHandler.SetFilePath("logs")
streamHandler.SetFileName("log")
```

Finally, create a new `Logger` object and add the desired handlers:

```go
logger := logger.NewLogger()
logger.AddHandler(consoleHandler)
logger.AddHandler(streamHandler)
```

You can now use the `Logger` object to log messages at any of the supported log levels:

```go
logger.Debug("This is a debug message")
logger.Info("This is an info message")
logger.Warning("This is a warning message")
logger.Error("This is an error message")
logger.Critical("This is a critical message")
```

## Customizing Output Format

The following placeholders are available for customizing the output format:

* `%d`: the current date and time (in the format "2006-01-02 15:04:05")
* `%l`: the log level
* `%m`: the message

For example, to include only the log level and message in the output, you could use the following format string:

```go
lineFormaterConsole.SetFormat("%l - %m")
```

## Supported Log Levels

The following log levels are supported:

* `DEBUG`
* `INFO`
* `WARNING`
* `ERROR`
* `CRITICAL`

## Multiple Handlers

It is possible to attach multiple handlers to a single `Logger` object. Each handler can have its own log level, formatter, and output destination. For example, you might want to send all debug and informational messages to the console, while sending only errors and critical messages to a log file. To do this, simply create additional handlers and add them to the logger:

```go
fileHandler := handler.NewStreamHandler()
fileHandler.SetFormater(lineFormaterFile)
fileHandler.SetLevel(levels.ERROR)
fileHandler.SetFilePath("logs")
fileHandler.SetFileName("error_log")

logger.AddHandler(fileHandler)
```

Now, both the console and the error log file will receive messages from the logger.

## Automatic Directory Creation

If the specified log file directory does not already exist, it will be automatically created when the `StreamHandler` writes to the file.
