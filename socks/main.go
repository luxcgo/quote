package main

import (
	"fmt"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/proxy"
)

func main() {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:7448", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	if resp, err := httpClient.Get("https://www.baidu.com"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}
