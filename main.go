package main

import (
	"os"
	"os/exec"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		return
	}

	if len(args) == 2 {
		cmd := exec.Command(args[1])
		cmd.Start()
		return
	}

	cmd := exec.Command(args[1], args[2:]...)
	cmd.Start()
}
