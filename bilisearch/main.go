package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	req, err := http.NewRequest("GET", "http://api.bilibili.com/x/web-interface/search/all/v2", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("keyword", "[2022-03-19 22-45-44][李壹壹]")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	req.AddCookie(&http.Cookie{Name: "SESSDATA", Value: "77500a0a%2C1661503483%2C6a09f021"})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	respBody := string(content)
	f, _ := os.Create("b.json")
	f.WriteString(respBody)
}
