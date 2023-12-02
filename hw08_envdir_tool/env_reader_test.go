package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/dchest/safefile"
	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("reading file from dir", func(t *testing.T) {

		err := os.Mkdir("/tmp/testdir", 0755)
		defer func() { _ = os.RemoveAll("/tmp/testdir") }()
		if err != nil {
			fmt.Println(err)
		}

		tmpfile, err := safefile.Create("/tmp/testdir/S.txt", 0755)
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

		var actual Environment
		actual, err = ReadDir("/tmp/testdir")
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range actual {
			actual[k] = v
		}
		var expected Environment
		expected = map[string]string{
			"S.txt": "temporary",
		}
		result := reflect.DeepEqual(expected, actual)
		require.True(t, result)
	})

	t.Run("reading invalid file", func(t *testing.T) {

		err := os.Mkdir("/tmp/testdir", 0755)
		defer func() { _ = os.RemoveAll("/tmp/testdir") }()
		if err != nil {
			fmt.Println(err)
		}
		tmpfile, err := safefile.Create("/tmp/testdir/S=q.txt", 0755)
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

		var actual Environment
		actual, err = ReadDir("/tmp/testdir")
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range actual {
			actual[k] = v
		}
		var expected Environment
		expected = map[string]string{}

		result := reflect.DeepEqual(expected, actual)
		require.True(t, result)
	})

	t.Run("reading multiline file", func(t *testing.T) {

		err := os.Mkdir("/tmp/testdir", 0755)
		defer func() { _ = os.RemoveAll("/tmp/testdir") }()
		if err != nil {
			fmt.Println(err)
		}

		tmpfile, err := safefile.Create("/tmp/testdir/file.txt", 0755)
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

		var actual Environment
		actual, err = ReadDir("/tmp/testdir")
		if err != nil {
			fmt.Println(err)
		}
		for k, v := range actual {
			actual[k] = v
		}
		var expected Environment
		expected = map[string]string{
			"file.txt": "the first line",
		}
		result := reflect.DeepEqual(expected, actual)
		require.True(t, result)
	})

}
