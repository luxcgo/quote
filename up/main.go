package main

import (
	"log"
	"path/filepath"
)

// "github.com/XiaoMiku01/biliup-go/login"
type Signed interface {
	~int
}

type myInt = int

func main() {
	move("archive")
}

var Filepath = "/downloads/bilibili/Zelo-Balance"

func move(dest string) {
	log.Println(filepath.Abs(dest))
	dir := (filepath.Dir(Filepath))
	log.Println(filepath.Join(dir, "archive"))
	base := filepath.Base(Filepath)
	dest = filepath.Join(dest, base)

	log.Println(base)
	log.Println(dest)
}
