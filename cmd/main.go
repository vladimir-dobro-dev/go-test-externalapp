package main

import (
	"fmt"

	"github.com/vladimir-dobro-dev/go-test-externalapp/logger"
)

func main() {
	fmt.Print("Run extarnal app, write log")
	logger.WriteLog()
}
