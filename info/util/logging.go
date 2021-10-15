package util

import (
	"fmt"
	"os"
	"time"
)

const (
        infoColor    = "\033[1;34m%s\033[0m"
        textColor    = "\033[1;36m%s\033[0m"
        warningColor = "\033[1;33m%s\033[0m"
        errorColor   = "\033[1;31m%s\033[0m"
        debugColor   = "\033[0;36m%s\033[0m"
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

	fmt.Printf("%s %s %s: %s\n", currentTime.Format("01-02-2006 15:04:05"), fmt.Sprintf(infoColor, " INFO"), l.Name, formatText(input))
}

func (l *Logger) Warn(format string, a ...interface {}) {
	var input string = fmt.Sprintf(format, a...)
	currentTime := time.Now()

	fmt.Printf("%s %s %s: %s\n", currentTime.Format("01-02-2006 15:04:05"), fmt.Sprintf(warningColor, " WARN"), l.Name, formatText(input))
}

func (l *Logger) Error(format string, a ...interface {}) {
	var input string = fmt.Sprintf(format, a...)
	currentTime := time.Now()

	fmt.Printf("%s %s %s: %s\n", currentTime.Format("01-02-2006 15:04:05"), fmt.Sprintf(errorColor, "ERROR"), l.Name, formatText(input))
}

func (l *Logger) Fatal(format string, a ...interface {}) {
	var input string = fmt.Sprintf(format, a...)
	currentTime := time.Now()

	fmt.Printf("%s %s %s: %s\n", currentTime.Format("01-02-2006 15:04:05"), fmt.Sprintf(errorColor, "FATAL"), l.Name, formatText(input))

	os.Exit(1)
}

func formatText(input string) (string) {
	return fmt.Sprintf(textColor, input)
}