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
	// fmt.Printf("\n***ENV EPNPT: (%v)***\n", env["EMPTY"])
	// fmt.Printf("\n***ENV UNSET: (%v)***\n", env["EMPTY"])
	// fmt.Printf("\nSTART")
	// fmt.Printf("===========================\n")
	// fmt.Printf("ENV: %v\n", env)
	// fmt.Printf("\n===========================")
	// fmt.Printf("============================END\n")
	RunCmd(command, env)
}
