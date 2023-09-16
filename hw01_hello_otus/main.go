package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	str := "Hello, OTUS!"
	revString := stringutil.Reverse(str)
	fmt.Println(revString)
}
