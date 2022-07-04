package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"unicode/utf8"
)

func main1() {
	f, _ := os.Open("a.html")
	b, _ := io.ReadAll(f)
	content := string(b)
	log.Println(string(b))

	title, _ := Match(`title="([^"]+)" target="_blank"`, content)
	roomName, _ := Match(`title="([^"]+)" class="router-link-exact-active`, content)
	url, _ := Match(`,"url":"([^"]+)"`, content)
	println(title, roomName, url)
}

func Match(pattern, content string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	submatch := re.FindAllStringSubmatch(content, -1)
	res := make([]string, 0)
	for _, v := range submatch {
		res = append(res, string(v[1]))
	}
	if len(res) < 1 {
		return "", errors.New("pattern not found")
	}
	return res[0], nil
}

func main() {
	u := "https:\u002F\u002Fali.pull.yximgs.com\u002Fgifkwai\u002FU5dJ_fCSBWo_ma1500.flv?auth_key=1651768905-0-0-eaabea389577e748649f94be8be68654&tsc=cdn&fd=1&ss=s1"
	utf8.DecodeRuneInString(src)
	fmt.Print(u)
}
