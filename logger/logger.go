package logger

import (
	"fmt"
	"os"
)

func WriteLog() {
	var file_path string = "D:/app.log"
	file, err := os.Create(file_path)
	if err != nil {
		fmt.Print("Unable to open file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString("1")
}
