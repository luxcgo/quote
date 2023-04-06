package main

import (
	"context"

	"github.com/go-olive/tv"
)

func main() {
	t, err := tv.New("bilibili", "8725120")
	if err != nil {
		return
	}
	context.WithValue()

}
