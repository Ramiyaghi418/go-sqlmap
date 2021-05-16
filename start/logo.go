package start

import (
	"fmt"
	"go-sqlmap/log"
)

// PrintLogo 打印Logo
func PrintLogo(version string, author string, url string) {
	logo := fmt.Sprintf(" _________      .__                         \n / "+
		"  _____/ _____|  |   _____ _____  ______  \n \\_____  \\ / ____"+
		"/  |  /     \\\\__  \\ \\____ \\ \n /        < <_|  |  |_|  Y Y"+
		"  \\/ __ \\|  |_> >\n/_______  /\\__   |____/__|_|  (____  /   _"+
		"_/ \n        \\/    |__|          \\/     \\/|__|           \n"+
		"Sqlmap-Simplified-Version-%s(golang)\nAuthor:%s "+
		"Github:%s", version, author, url)
	fmt.Println(logo)
	log.Info("start go-sqlmap...")
}
