package log

import (
	"fmt"
	"time"
)

// Debug 日志
func Debug(log string) {
	fmt.Printf("[DEBUG] [%s] %s\n", getTime(), log)
}

// Info 日志
func Info(log string) {
	fmt.Printf("[INFO] [%s] %s\n", getTime(), log)
}

// Warn 日志
func Warn(log string) {
	fmt.Printf("[WARN] [%s] %s\n", getTime(), log)
}

// Error 日志
func Error(log string) {
	fmt.Printf("[ERROR] [%s] %s\n", getTime(), log)
}

// Fatal 日志
func Fatal(log string) {
	fmt.Printf("[FATAL] [%s] %s\n", getTime(), log)
}

// Panic 日志
func Panic(log string) {
	fmt.Printf("[PANIC] [%s] %s\n", getTime(), log)
}

// 获取当前时间
func getTime() string {
	currentTime := time.Now().Format("15:04:05")
	return currentTime
}
