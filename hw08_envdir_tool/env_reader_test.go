package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/dchest/safefile"
	"github.com/stretchr/testify/require"
)

func TestWithDir(t *testing.T) {
	t.Run("reading file from dir", func(t *testing.T) {
		err := os.Mkdir("/tmp/testdir", 0o755)
		defer func() { _ = os.RemoveAll("/tmp/testdir") }()
		if err != nil {
			fmt.Println(err)
		}

		tmpfile, err := safefile.Create("/tmp/testdir/S.txt", 0o755)
		if err != nil {
			fmt.Println(err)
		}
		defer tmpfile.Close()

		content := []byte("temporary")
		if _, err := tmpfile.Write(content); err != nil {
			log.Fatal(err)
		}

		if err := tmpfile.Commit(); err != nil {
			log.Fatal(err)
		}

		var result Environment
		result, err = ReadDir("/tmp/testdir")
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range result {
			result[k] = v
		}
		expected := Environment{
			"S.txt": "temporary",
		}
		require.Equal(t, expected, result)
	})
}

func TestWithFiles(t *testing.T) {
	t.Run("reading invalid file", func(t *testing.T) {
		err := os.Mkdir("/tmp/testdir", 0o755)
		defer func() { _ = os.RemoveAll("/tmp/testdir") }()
		if err != nil {
			fmt.Println(err)
		}
		tmpfile, err := safefile.Create("/tmp/testdir/S=q.txt", 0o755)
		if err != nil {
			fmt.Println(err)
		}
		defer tmpfile.Close()

		content := []byte("temporary")
		if _, err := tmpfile.Write(content); err != nil {
			log.Fatal(err)
		}

		if err := tmpfile.Commit(); err != nil {
			log.Fatal(err)
		}

		var result Environment
		result, err = ReadDir("/tmp/testdir")
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range result {
			result[k] = v
		}
		expected := Environment{}

		require.Equal(t, expected, result)
	})

	t.Run("reading multiline file", func(t *testing.T) {
		err := os.Mkdir("/tmp/testdir", 0o755)
		defer func() { _ = os.RemoveAll("/tmp/testdir") }()
		if err != nil {
			fmt.Println(err)
		}

		tmpfile, err := safefile.Create("/tmp/testdir/file.txt", 0o755)
		if err != nil {
			fmt.Println(err)
		}
		defer func() { _ = tmpfile.Close() }()

		content1 := []byte("the first line\nthe second line\nthe third line")
		if _, err := tmpfile.Write(content1); err != nil {
			log.Fatal(err)
		}

		if err := tmpfile.Commit(); err != nil {
			log.Fatal(err)
		}

		var result Environment
		result, err = ReadDir("/tmp/testdir")
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range result {
			result[k] = v
		}
		expected := Environment{
			"file.txt": "the first line",
		}
		require.Equal(t, expected, result)
	})
}
