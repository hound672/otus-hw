package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	path := os.Args[1]
	command := os.Args[2:]

	env, err := ReadDir(path)
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		log.Fatalf("Err ReadEnv: %v", err)
	}

	RunCmd(command, env)
}
