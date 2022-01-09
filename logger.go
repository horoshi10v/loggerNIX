package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	TagInfo    = "Info: "
	TagError   = "Error: "
	TagWarning = "Warning: "
	TagDebug   = "Debug: "
	timeFormat = "15:04:05.000"
	//FlagsTime  = time.Second | time.Millisecond | time.Microsecond
)

type logger struct {
	logPath string
}

func NewLogger(path string) logger {
	return logger{logPath: path}
}

func LoggInFile(logPath, level, msg, time string) {
	buff := []string{
		level,
		msg,
		time,
	}
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	writer := bufio.NewWriter(lf)
	if err != nil {
		fmt.Println("Unable to create/open file:", err)
		os.Exit(1)
	}
	defer func(lf *os.File) {
		err := lf.Close()
		if err != nil {
			fmt.Println("Unable to close file:", err)
			os.Exit(1)
		}
	}(lf)
	for _, row := range buff {
		writer.WriteString(row)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func PrintMsg(tag, message string) {
	timeF := time.Now().Format(timeFormat)
	fmt.Printf("%v \nMessage: %v \nTime:%v \n",
		tag, message, timeF)
}
func (l *logger) Info(message string) {
	PrintMsg(TagInfo, message)
}
func (l *logger) InfoF(message string) {
	timeF := time.Now().Format(timeFormat)
	LoggInFile(l.logPath, TagInfo, message, timeF)
}
func (l *logger) Error(message string) {
	PrintMsg(TagError, message)
}
func (l *logger) ErrorF(message string) {
	timeF := time.Now().Format(timeFormat)
	LoggInFile(l.logPath, TagError, message, timeF)
}
func (l *logger) Warning(message string) {
	PrintMsg(TagWarning, message)
}
func (l *logger) WarningF(message string) {
	timeF := time.Now().Format(timeFormat)
	LoggInFile(l.logPath, TagWarning, message, timeF)
}
func (l *logger) Debug(message string) {
	PrintMsg(TagDebug, message)
}
func (l *logger) DebugF(message string) {
	timeF := time.Now().Format(timeFormat)
	LoggInFile(l.logPath, TagDebug, message, timeF)
}
