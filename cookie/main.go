package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main1() {

	resp, _ := http.Get("https://www.kuaishou.com/")
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "did" {
			fmt.Println("Found a cookie named:", cookie.Name, cookie.Value)
		}
	}
}
func main2() {
	client := http.Client{}
	url := "http://live.kuaishou.com/"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("referer", "http://live.kuaishou.com/")
	resp, err := client.Do(req) //send request
	if err != nil {
		return
	}
	for _, cookie := range resp.Cookies() {
		fmt.Println("Found a cookie named:", cookie.Name, cookie.Value)
	}
	// log.Println(resp.Header)
}

func main3() {
	url := "https://live.kuaishou.com/profile/3xhmaesbuemazv9"
	f := NewFetcherHttps("")
	resp, body, err := f.Get(url)
	println("status:", resp.StatusCode)
	// println("body:", string(body))
	println(err)
	for _, v := range f.Cookies {
		println(v.Name, v.Value)
	}
	log.Println(string(body))
}

// __ac_signature _02B4Z6wo00f01yn2FbAAAIDAvSBOvcltnpsp1hEAAKggaTgSUIMUGHQ5S2Y.hlSDaNTYvBxaKDqUK0QzpehWK2a0FwmW18CBePweLyQfAwXBusx6p6DEOLfbYB6DqcegwE4yFfsYtEzJ6tnW32
// __ac_nonce  0626ba34c0050c318cee7

// curl 'https://www.douyin.com/' \
//   -H 'authority: www.douyin.com' \
//   -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
//   -H 'accept-language: zh-CN,zh;q=0.9' \
//   -H 'cache-control: max-age=0' \
//   -H 'cookie: ttwid=1%7CxLR3b0G-jsszeUrhlZ3Zwah0fGuyvgJJ_VgLv3sPsH4%7C1651216801%7C2c10c2ead47cc638aa3e4cef46aa804e8b8ab8e4b70dc14a0b9a9a028d4eeaea; _tea_utm_cache_6383=undefined; douyin.com; strategyABtestKey=1651216805.54; passport_csrf_token=50324329f7e7076669dc4845805f70e8; passport_csrf_token_default=50324329f7e7076669dc4845805f70e8; _tea_utm_cache_1300=undefined; s_v_web_id=verify_l2k3utj5_4O4KAbmC_FTSG_4FCs_8r4x_vOqoSRT6zIAk; ttcid=58d6e8b383c043948b11a43e18aca4e126; THEME_STAY_TIME=299515; IS_HIDE_THEME_CHANGE=1; pwa_guide_count=3; live_can_add_dy_2_desktop=0; odin_tt=51b7dd0c53fb2b19f422af9715f94db8cd51bc297720b7388f2cc87a3f39ccd4f66f34023f63b83cdc985847cc5448ccaeeb2686ed5ffefb8555bbda37827303926ba6a523ff407b54193dc4133c2141; _tea_utm_cache_2285=undefined; __ac_nonce=0626baae700aa6ca9c2ac; __ac_signature=_02B4Z6wo00f0126u3.AAAIDA-niE.TU5N4dujttAALnLO4QDrAxWB90ASy.c0GLLqQCZlXZIQ6xZ1.DvbLH69nTQ6ROB0INRZ.ftzOcQSnuAWKyhUG5GKrTR8G6OL-uejbiArgs1CO5GNucN4d; home_can_add_dy_2_desktop=1; msToken=sgB7qgEYrtFfqHAzCMHD_uUfbYNUth1nQRVqwvwdHQ1YbB59j_e59H3kCWcj0FOr1ROb5myiaPfMnMU64OnRP1LDZgu02VPF7S6laWjy-gKmIf2L4G9vGAJZuc36J8cD; msToken=sQ4OMLEMIhKX2VyodiL-qkKukZya_cX6KZAzk_8nrcpsPMS2kwXuV0uSR9M8N-MMLTQ2QDB3VB8gq3kVHzw-iYiIVkyVzOVIvxjv-x3-dY9rlZB2_PPB8Aw=; tt_scid=HZO9O-2mtDN6XpLnbZCmmxQyAOHLqqWEzgDw1KzSNfH3WAk2SXkARihykKpzoIfwbcf6' \
//   -H 'referer: https://www.douyin.com/' \
//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'sec-ch-ua-platform: "macOS"' \
//   -H 'sec-fetch-dest: document' \
//   -H 'sec-fetch-mode: navigate' \
//   -H 'sec-fetch-site: same-origin' \
//   -H 'sec-fetch-user: ?1' \
//   -H 'upgrade-insecure-requests: 1' \
//   -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36' \
//   --compressed

func main4() {
	f := NewFetcherHttps("live.kuaishou.com")

	resp, body, err := f.Get("/")
	if err != nil {
		return
	}
	resp, body, err = f.Get("/profile/3xhmaesbuemazv9")

	println("status:", resp.StatusCode)
	println("body:", string(body))
	ff, _ := os.Create("a.html")
	ff.WriteString(string(body))
}

func main() {
	req, _ := http.NewRequest("GET", "https://live.kuaishou.com/profile/xiatian565", nil)
	req.Header.Add("cookie", "did=web_d079abeeeba77349c6eb5724363b8958")
	resp, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(resp.Body)
	println("status:", resp.StatusCode)
	println("body:", string(body))
	ff, _ := os.Create("a.html")
	ff.WriteString(string(body))
}

func main6() {
	u := "https:\u002F\u002Fali-origin.pull.yximgs.com\u002Fgifshow\u002F76v0SFKHkyQ_ma1500.flv?auth_key=1651763572-0-0-8ae3cae2f2863da4390390798b27d37d&tsc=origin&oidc=watchmen&sidc=5242&fd=1&ss=s20"
	log.Println(u)
}
