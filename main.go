package main

import (
	"encoding/json"
	"os"
)

func main() {
	res, _ := json.Marshal(&Point{X: 1})
	probJson, _ := os.Create("problems.json")
	probJson.Write(res)
}

type Point struct {
	X int
}
