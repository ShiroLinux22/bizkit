/*
	Struct(s) for consistent logging
    Copyright (C) 2021 jacany <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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