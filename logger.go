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
func Info(message string) {
	PrintMsg(TagInfo, message)
}
func InfoF(message, path string) {
	timeF := time.Now().Format(timeFormat)
	LoggInFile(path, TagInfo, message, timeF)
}
func Error(message string) {
	PrintMsg(TagError, message)
}
func ErrorF(message, path string) {
	timeF := time.Now().Format(timeFormat)
	LoggInFile(path, TagError, message, timeF)
}
func Warning(message string) {
	PrintMsg(TagWarning, message)
}
func WarningF(message, path string) {
	timeF := time.Now().Format(timeFormat)
	LoggInFile(path, TagWarning, message, timeF)
}
func Debug(message string) {
	PrintMsg(TagDebug, message)
}
func DebugF(message, path string) {
	timeF := time.Now().Format(timeFormat)
	LoggInFile(path, TagDebug, message, timeF)
}
