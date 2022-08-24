package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/levigross/grequests"
)

var (
	hostType = "apiv1"
	ro       = &grequests.RequestOptions{
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36",
		RequestTimeout: 10 * time.Second,
		Headers: map[string]string{
			"Connection":      "keep-alive",
			"Accept":          "*/*",
			"Accept-Encoding": "*",
			"Accept-Language": "zh-CN,zh;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5",
		},
	}
)

func main() {
	GenTsUrls()
}
func GenTsUrls() {
	m3u8Url := "https://d3vd9lfkzbru3h.cloudfront.net/b895f864fec797f9a481_domado0129_39585681303_1660997202/chunked/index-dvr.m3u8"
	m3u8Host := getHost(m3u8Url, hostType)
	m3u8Body, _ := os.ReadFile("b.m3u8")
	ts_list := getTsList(m3u8Host, string(m3u8Body))
	fmt.Println("待下载 ts 文件数量:", len(ts_list))

	f, err := os.Create("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, v := range ts_list {
		s := fmt.Sprintf("%s\n", v.Url)
		if strings.Contains(s, "unmuted") {
			continue
		}
		f.WriteString(s)
	}
	return
}

func getHost(Url, ht string) (host string) {
	u, err := url.Parse(Url)
	if err != nil {
		log.Fatal(err)
	}
	switch ht {
	case "apiv1":
		host = u.Scheme + "://" + u.Host + filepath.Dir(u.EscapedPath())
	case "apiv2":
		host = u.Scheme + "://" + u.Host
	}
	return
}

func getM3u8Body(Url string) string {
	r, err := grequests.Get(Url, ro)
	if err != nil {
		log.Fatal(err)
	}
	return r.String()
}

func getTsList(host, body string) (tsList []TsInfo) {
	lines := strings.Split(body, "\n")
	index := 0
	var ts TsInfo
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") && line != "" {
			//有可能出现的二级嵌套格式的m3u8,请自行转换！
			index++
			if strings.HasPrefix(line, "http") {
				ts = TsInfo{
					Url: line,
				}
				tsList = append(tsList, ts)
			} else {
				ts = TsInfo{
					Url: fmt.Sprintf("%s/%s", host, line),
				}
				tsList = append(tsList, ts)
			}
		}
	}
	return
}

type TsInfo struct {
	Url string
}
