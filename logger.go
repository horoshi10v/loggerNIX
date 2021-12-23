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

func PrintMsg(tag, message string, toFile bool) {
	timeF := time.Now().Format("15:04:05.000")
	if !toFile {
		fmt.Printf("%v \nMessage: %v \n Time:%v \n",
			tag, message, timeF)
	} else {
		var logPath string
		LoggInFile(logPath, tag, message, timeF)
	}
}

func Info(message string, toFile bool) {
	tag := TagInfo
	PrintMsg(tag, message, toFile)
}

func Error(message string, toFile bool) {
	tag := TagError
	PrintMsg(tag, message, toFile)
}

func Warning(message string, toFile bool) {
	tag := TagWarning
	PrintMsg(tag, message, toFile)
}

func Debug(message string, toFile bool) {
	tag := TagDebug
	PrintMsg(tag, message, toFile)
}
