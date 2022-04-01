package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	logrus "github.com/sirupsen/logrus"
)

var httpTr = &http.Transport{
	//控制主机的最大空闲连接数，0为没有限制
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
	//长连接在关闭之前，保持空闲的最长时间，0表示没限制。
	IdleConnTimeout: 60 * time.Second,
}

var client = &http.Client{
	Timeout:   20 * time.Second,
	Transport: httpTr,
}

func init() {
	formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	logrus.SetFormatter(&formatter)
}

func SendRequest(url string, method string, contentType string, param io.Reader, respData interface{}, headers ...map[string]string) error {
	var err error
	req, err := http.NewRequest(method, url, param)
	if err != nil {
		err = fmt.Errorf("create http request failed: %s", err.Error())
		return err
	}

	req.Header.Set("Content-Type", contentType)

	headers = []map[string]string{{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 Edg/94.0.992.38",
		"referer":    "https://live.douyin.com/",
		"cookie":     "__ac_nonce=0623f2d6c00831bba17e4; __ac_signature=_02B4Z6wo00f01jp4M0wAAIDDQrbYS7tdHpY6WDfAAOyva7;",
	}}

	if headers != nil {
		for _, h := range headers {
			for k, v := range h {
				req.Header.Set(k, v)
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("send http request failed: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read response data failed: %s", err.Error())
		return err
	}
	logrus.Debug("read the response ok:", string(respBytes))

	// 解析响应
	// if err = json.Unmarshal(respBytes, respData); err != nil {
	// 	logrus.Error(err.Error())
	// 	err = fmt.Errorf("unmarshal response data failed: %s", err.Error())
	// 	return err
	// }

	log.Println(string(respBytes))
	return nil
}

func PostFormDataRequest(urlstr string, requestData map[string]string, responseData interface{}) error {
	var err error
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if requestData != nil {
		for k, v := range requestData {
			err = writer.WriteField(k, v)
			if err != nil {
				err = fmt.Errorf("write form field %s:%s failed: %s", k, v, err.Error())
				return err
			}
			logrus.Debug("build form data ok.")
		}
	}
	if err = writer.Close(); err != nil {
		err = fmt.Errorf("close multipart form writer failed: %s", err.Error())
	}

	err = SendRequest(urlstr, "POST", writer.FormDataContentType(), body, responseData)
	if err != nil {
		return fmt.Errorf("do sendRequest error with param %s:%s", body, err.Error())
	}

	return nil
}

func main() {
	roomInit := new(RoomInit)
	PostFormDataRequest("https://api.live.bilibili.com/room/v1/Room/room_init",
		map[string]string{
			"id": "764155",
		},
		roomInit)
	log.Printf("%+v", roomInit)
}

type RoomInit struct {
	Code int64 `json:"code"`
	Data struct {
		RoomID     int64 `json:"room_id"`
		LiveStatus int64 `json:"live_status"`
	}
}

func main2() {
	var roomInit interface{}
	PostFormDataRequest("https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo",
		map[string]string{
			"room_id":  "81711",
			"protocol": "0,1",
			"format":   "0,1,2",
			"codec":    "0,1",
			"qn":       "10000",
			"platform": "h5",
			"ptype":    "8",
		},
		roomInit)
	log.Printf("%+v", roomInit)
}

func main3() {
	var roomInit interface{}
	PostFormDataRequest("https://live.douyin.com/235808786065",
		map[string]string{
			// "room_id":  "81711",
			// "protocol": "0,1",
			// "format":   "0,1,2",
			// "codec":    "0,1",
			// "qn":       "10000",
			// "platform": "h5",
			// "ptype":    "8",
		},
		roomInit)
	log.Printf("%+v", roomInit)
}
