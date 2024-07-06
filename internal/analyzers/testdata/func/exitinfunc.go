package main

import (
	"os"
)

func main() {
	println("Hello, World!")
	Exit(1)
}

func Exit(code int) {
	os.Exit(code)
}
