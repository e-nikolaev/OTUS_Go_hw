package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const STRING = "Hello, OTUS!"

func main() {
	fmt.Println(stringutil.Reverse(STRING))
}
