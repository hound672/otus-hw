package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	str := "Hello, OTUS!"
	revstring := stringutil.Reverse(str)
	fmt.Println(revstring)
}
