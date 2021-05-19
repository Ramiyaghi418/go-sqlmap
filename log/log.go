package log

import (
	"fmt"
	"time"
)

// Info 日志
func Info(log string) {
	fmt.Printf("[INFO] [%s] %s\n", getTime(), log)
}

// InfoLine 不含换行日志
func InfoLine(log string) {
	fmt.Printf("[INFO] [%s] %s", getTime(), log)
}

// Error 日志
func Error(log string) {
	fmt.Printf("[ERROR] [%s] %s\n", getTime(), log)
}

// 获取当前时间
func getTime() string {
	currentTime := time.Now().Format("15:04:05")
	return currentTime
}
