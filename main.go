package main

import (
	"os"

	"kusionstack.io/kpt-kcl-sdk/pkg/runner"
)

func main() {
	if err := runner.Run(); err != nil {
		os.Exit(1)
	}
}
