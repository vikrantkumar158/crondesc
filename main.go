package main

import (
	"fmt"
	"os"

	"crondesc/logic"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Error: Please provide a valid cron string")
		os.Exit(1)
	}

	cronString := os.Args[1]
	logic.ProcessCronString(cronString)
}
