package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var env Environment
		env = map[string]string{
			"STR1": "foo",
			"STR2": "bar",
		}
		cmd := []string{"bash", "-c", "echo $STR1$STR2"}
		code := RunCmd(cmd, env)

		fmt.Println(cmd)
		require.Equal(t, 0, code)
	})

	t.Run("Invalid command error", func(t *testing.T) {
		var env Environment
		env = map[string]string{}
		cmd := []string{"bash", "-c", "ls /xxx"}
		code := RunCmd(cmd, env)

		fmt.Println(cmd)
		require.Equal(t, 1, code)
	})
}
