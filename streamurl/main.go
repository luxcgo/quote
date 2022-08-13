package main

import (
	"os"
	"time"

	"github.com/go-olive/tv"
)

func main() {
	t, err := tv.New("bilibili", "8725120")
	if err != nil {
		return
	}

	f, _ := os.Create("burl.txt")
	defer f.Close()
	for i := 0; i < 50; i++ {
		t.Snap()
		u, _ := t.StreamUrl()
		f.WriteString(u)
		f.WriteString("\n")
		time.Sleep(time.Millisecond * 50)
	}
}
