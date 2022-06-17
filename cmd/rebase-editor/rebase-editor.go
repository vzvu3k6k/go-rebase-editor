package main

import (
	"fmt"
	"os"

	"github.com/vzvu3k6k/go-rebase-editor"
)

func main() {
	if err := rebase.Run(os.Args[1:]); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
