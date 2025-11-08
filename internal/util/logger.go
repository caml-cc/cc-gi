package util

import (
	"fmt"
	"time"
)

var isVerbose bool

func SetVerbose(verbose bool) {
	isVerbose = verbose
}

func now() string {
	return time.Now().Format("15:04:05")
}

func DebugLog(message string, args ...any) {
	if isVerbose {
		fmt.Printf(ColorBlue+now()+"[DEBUG] "+ColorReset+message+"\n", args...)
	}
}

func ErrorLog(message string, args ...any) {
	if isVerbose {
		fmt.Printf(ColorRed+now()+"[ERROR] "+ColorReset+message+"\n", args...)
	}
}

func InfoLog(message string, args ...any) {
	if isVerbose {
		fmt.Printf(ColorGreen+now()+"[INFO] "+ColorReset+message+"\n", args...)
	}
}

func WarnLog(message string, args ...any) {
	if isVerbose {
		fmt.Printf(ColorYellow+now()+"[WARN] "+ColorReset+message+"\n", args...)
	}
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorGray   = "\033[90m"
)
