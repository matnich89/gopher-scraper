package logging

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

var (
	infoColor    = color.New(color.FgCyan)
	successColor = color.New(color.FgGreen)
	errorColor   = color.New(color.FgRed)
	warnColor    = color.New(color.FgYellow)
	debugColor   = color.New(color.FgWhite)
)

func LogInfo(format string, a ...interface{}) {
	infoColor.Printf("[%s] ℹ️  INFO: %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(format, a...))
}

func LogSuccess(format string, a ...interface{}) {
	successColor.Printf("[%s] ✅ SUCCESS: %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(format, a...))
}

func LogError(format string, a ...interface{}) {
	errorColor.Printf("[%s] ❌ ERROR: %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(format, a...))
}

func LogWarning(format string, a ...interface{}) {
	warnColor.Printf("[%s] ⚠️  WARNING: %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(format, a...))
}

func LogDebug(format string, a ...interface{}) {
	debugColor.Printf("[%s] 🔍 DEBUG: %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(format, a...))
}
