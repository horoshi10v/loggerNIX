package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	TagInfo    = "INFO:\t"
	TagError   = "ERROR:\t"
	TagWarning = "WARNING:\t"
	TagDebug   = "DEBUG:\t"
	timeFormat = time.RFC1123
)

var timeF = time.Now().Format(timeFormat)

type Logger struct {
	FilePath string
}

func NewLogger(File string) Logger {
	return Logger{FilePath: File}
}

func (l *Logger) PrintMsg(tag, message string) {
	msg := timeF + " " + tag + " " + message + "\n"
	fmt.Printf(msg)
}

func (l *Logger) InfoFile(message string) {
	LoggInFile(l.FilePath, TagInfo, message)
}
func (l *Logger) Info(message string) {
	l.PrintMsg(TagInfo, message)
}
func (l *Logger) Error(message string) {
	l.PrintMsg(TagError, message)
}
func (l *Logger) ErrorFile(message string) {
	LoggInFile(l.FilePath, TagError, message)
}
func (l *Logger) Warning(message string) {
	l.PrintMsg(TagWarning, message)
}
func (l *Logger) WarningFile(message string) {
	LoggInFile(l.FilePath, TagWarning, message)
}
func (l *Logger) Debug(message string) {
	l.PrintMsg(TagDebug, message)
}
func (l *Logger) DebugFile(message string) {
	LoggInFile(l.FilePath, TagDebug, message)
}

func LoggInFile(logPath, level, msg string) {
	buff := []string{
		level,
		msg,
		timeF,
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
