package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main1() {
	url := "https://www.twitch.tv/videos/156447"
	for i := 5000; i < 7000; i++ {
		log.Println(i)
		newUrl := url + strconv.Itoa(i)
		content, _ := GetURLContent(newUrl)
		content = strings.ToLower(content)
		if strings.Contains(content, "domado0129") {
			log.Println(newUrl)
			break
		}
		time.Sleep(time.Millisecond)
	}

	// content, _ := GetURLContent("https://www.twitch.tv/videos/1566372736")

	// log.Println(content)
}

func main() {
	url := "https://www.twitch.tv/videos/156637273"
	for i := 0; i < 7000; i++ {
		log.Println(i)
		newUrl := url + strconv.Itoa(i)
		content, _ := GetURLContent(newUrl)
		content = strings.ToLower(content)
		if strings.Contains(content, "domado0129") {
			log.Println(newUrl)
			break
		}
		time.Sleep(time.Millisecond)
	}

	// content, _ := GetURLContent("https://www.twitch.tv/videos/1566372736")

	// log.Println(content)
}

func GetURLContent(url string) (string, error) {
	webUserAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", webUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	content := string(respBody)
	return content, nil
}
