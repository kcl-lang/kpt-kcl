package main

import (
	"os"

	"kcl-lang.io/kpt-kcl/pkg/runner"
)

func main() {
	if err := runner.Run(); err != nil {
		os.Exit(1)
	}
}
