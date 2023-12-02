package main

import (
	"fmt"
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
		fmt.Println(err)
		_, _ = os.Stderr.WriteString(err.Error())
	}
	RunCmd(command, env)
}
