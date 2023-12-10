package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	// Place your code here
	list := make(map[string]string)

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "readDir failed")
	}

	for _, f := range files {
		name := f.Name()
		if !strings.Contains(name, "=") {
			file, err := os.OpenFile(dir+"/"+name, os.O_RDONLY, 0o755)
			if err != nil {
				return nil, fmt.Errorf("open file: %w", err)
			}
			defer func() { _ = file.Close() }()
			scanner := bufio.NewScanner(file)
			scanner.Scan()

			rawEnv := scanner.Bytes()
			envValue := string(bytes.ReplaceAll(rawEnv, []byte("\x00"), []byte("\n")))
			envValue = strings.TrimRight(envValue, " ")

			list[name] = envValue

			if err := scanner.Err(); err != nil {
				return nil, fmt.Errorf("scanner.Err: %w", err)
			}
		}
	}

	return list, nil
}
