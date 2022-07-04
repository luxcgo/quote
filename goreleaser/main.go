package main

import "os"

func main() {
	println(os.Getwd())
	println(os.Executable())
	println(os.Args[0])
}
