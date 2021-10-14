package util

import (
	"fmt"
	"os"
	"time"
)

const (
        InfoColor    = "\033[1;34m%s\033[0m"
        TextColor    = "\033[1;36m%s\033[0m"
        WarningColor = "\033[1;33m%s\033[0m"
        ErrorColor   = "\033[1;31m%s\033[0m"
        DebugColor   = "\033[0;36m%s\033[0m"
)

type Logger struct {
	Name string
}

func (l *Logger) WarnOnError(err error, msg string) {
	if err != nil {
		l.Warn("%s: %s", msg, err)
	}
}

func (l *Logger) FatalOnError(err error, msg string) {
	if err != nil {
		l.Fatal("%s: %s", msg, err)
	}
}

func (l *Logger) Info(format string, a ...interface {}) {
	var input string = fmt.Sprintf(format, a...)
	currentTime := time.Now()

	fmt.Printf("%s %s %s: %s\n", currentTime.Format("01-02-2006 15:04:05"), fmt.Sprintf(InfoColor, " INFO"), l.Name, formatText(input))
}

func (l *Logger) Warn(format string, a ...interface {}) {
	var input string = fmt.Sprintf(format, a...)
	currentTime := time.Now()

	fmt.Printf("%s %s %s: %s\n", currentTime.Format("01-02-2006 15:04:05"), fmt.Sprintf(WarningColor, " WARN"), l.Name, formatText(input))
}

func (l *Logger) Error(format string, a ...interface {}) {
	var input string = fmt.Sprintf(format, a...)
	currentTime := time.Now()

	fmt.Printf("%s %s %s: %s\n", currentTime.Format("01-02-2006 15:04:05"), fmt.Sprintf(ErrorColor, "ERROR"), l.Name, formatText(input))
}

func (l *Logger) Fatal(format string, a ...interface {}) {
	var input string = fmt.Sprintf(format, a...)
	currentTime := time.Now()

	fmt.Printf("%s %s %s: %s\n", currentTime.Format("01-02-2006 15:04:05"), fmt.Sprintf(ErrorColor, "FATAL"), l.Name, formatText(input))

	os.Exit(1)
}

func formatText(input string) (string) {
	return fmt.Sprintf(TextColor, input)
}