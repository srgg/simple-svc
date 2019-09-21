package main

import (
	"github.com/srgg/simple-svc/internal/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
