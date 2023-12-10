package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"testing"

	"github.com/dchest/safefile"
	"github.com/stretchr/testify/require"
)

func TestWithDir(t *testing.T) {
	t.Run("reading file from dir", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "*")
		if err != nil {
			panic(err)
		}

		defer func() { _ = os.RemoveAll(tmpDir) }()

		tmpFile, err := safefile.Create(path.Join(tmpDir, "S.txt"), 0o755)
		if err != nil {
			fmt.Println(err)
		}
		defer func() { _ = tmpFile.Close() }()

		content := []byte("temporary")
		if _, err := tmpFile.Write(content); err != nil {
			log.Fatal(err)
		}

		if err := tmpFile.Commit(); err != nil {
			log.Fatal(err)
		}

		var result Environment
		result, err = ReadDir(tmpDir)
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
		tmpDir, err := os.MkdirTemp("", "*")
		if err != nil {
			panic(err)
		}

		defer func() { _ = os.RemoveAll(tmpDir) }()

		tmpFile, err := safefile.Create(path.Join(tmpDir, "S=q.txt"), 0o755)
		if err != nil {
			fmt.Println(err)
		}
		defer func() { _ = tmpFile.Close() }()

		content := []byte("temporary")
		if _, err := tmpFile.Write(content); err != nil {
			log.Fatal(err)
		}

		if err := tmpFile.Commit(); err != nil {
			log.Fatal(err)
		}

		var result Environment
		result, err = ReadDir(tmpDir)
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
		tmpDir, err := os.MkdirTemp("", "*")
		if err != nil {
			panic(err)
		}

		defer func() { _ = os.RemoveAll(tmpDir) }()

		tmpFile, err := safefile.Create(path.Join(tmpDir, "file.txt"), 0o755)
		if err != nil {
			fmt.Println(err)
		}
		defer func() { _ = tmpFile.Close() }()

		content1 := []byte("the first line\nthe second line\nthe third line")
		if _, err := tmpFile.Write(content1); err != nil {
			log.Fatal(err)
		}

		if err := tmpFile.Commit(); err != nil {
			log.Fatal(err)
		}

		var result Environment
		result, err = ReadDir(tmpDir)
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
