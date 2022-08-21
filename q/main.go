package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// filepathStr := "/Users/lucas/github/quote/q"
	matches, err := filepath.Glob("[0-9]*.ts")
	fmt.Println(matches, err)
}

func change(s *[]int) {
	*s = append(*s, 1, 2, 3)
	(*s)[0] = -1
}
