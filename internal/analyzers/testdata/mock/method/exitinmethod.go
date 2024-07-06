package main

import "os"

func main() {
	println("Hello, World!")
	t := ExitStruct{}
	t.Exit(1)
}

type ExitStruct struct {
}

func (e ExitStruct) Exit(code int) {
	os.Exit(code)
}
