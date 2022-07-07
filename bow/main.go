package main

import (
	"fmt"
	"html"
	"log"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	u := "http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=3d&issueCount=30"

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	err := bow.Open(u)
	if err != nil {
		panic(err)
	}
	// Outputs: "reddit: the front page of the internet"
	fmt.Println(bow.Title())

	// Click the link for the newest submissions.
	// bow.Click("number u-pad10")
	resp := html.UnescapeString((bow.Body()))
	var auto TdAutoGenerated
	jsoniter.UnmarshalFromString(resp, &auto)
	for _, v := range auto.Result {
		println(v.Code)
		println(v.Date)
		log.Println(strings.Split(v.Red, ","))
		println(v.Red)
	}
}

// https://webapi.sporttery.cn/gateway/lottery/getHistoryPageListV1.qry?gameNo=35&provinceId=0&pageSize=30&isVerify=1&pageNo=1

func main2() {
	s := "0 3 4"
	log.Println(strings.Fields(s))
}

type TdAutoGenerated struct {
	State     int    `json:"state"`
	Message   string `json:"message"`
	PageCount int    `json:"pageCount"`
	CountNum  int    `json:"countNum"`
	Tflag     int    `json:"Tflag"`
	Result    []struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		DetailsLink string `json:"detailsLink"`
		VideoLink   string `json:"videoLink"`
		Date        string `json:"date"`
		Week        string `json:"week"`
		Red         string `json:"red"`
		Blue        string `json:"blue"`
		Blue2       string `json:"blue2"`
		Sales       string `json:"sales"`
		Poolmoney   string `json:"poolmoney"`
		Content     string `json:"content"`
		Addmoney    string `json:"addmoney"`
		Addmoney2   string `json:"addmoney2"`
		Msg         string `json:"msg"`
		Z2Add       string `json:"z2add"`
		M2Add       string `json:"m2add"`
		Prizegrades []struct {
			Type      int    `json:"type"`
			Typenum   string `json:"typenum"`
			Typemoney string `json:"typemoney"`
		} `json:"prizegrades"`
	} `json:"result"`
}