package main

import (
	"log"

	jsoniter "github.com/json-iterator/go"
)

func main1() {
	sm := NewRWMap[string, Show](2)
	s1 := Show{ID: "s1", Name: "s1name"}
	sm.Set(s1.ID, s1)

	b1 := &Bout{
		showMap: sm,
		showID:  s1.ID,
	}

	log.Println(b1.showID)
	b1s1, _ := b1.showMap.Get(b1.showID)
	log.Println(b1s1.Name)

	s1.Name = "changeds1name"
	sm.Set(s1.ID, s1)
	log.Println(b1.showID)
	b1s1, _ = b1.showMap.Get(b1.showID)
	log.Println(b1s1.Name)
}

type Show struct {
	ID   string
	Name string
}

type Bout struct {
	showMap *RWMap[string, Show]
	showID  string
	s       *Show
}

func (b *Bout) P() {
	log.Println(b.s.ID)
}

func main2() {
	s := &Show{ID: "1"}

	b := &Bout{s: s}
	b.P()

	newS := &Show{ID: "2"}
	*s = *newS
	b.P()
}

func main() {
	s :=
		`{
		"code": "0000",
		"data": {
			"PortalUsername": "olive",
			"PortalPassword": "olive",
			"SaveDir": "/",
			"OutTmpl": "[{{ .StreamerName }}][{{ .RoomName }}][{{ now | date \"2006-01-02 15-04-05\"}}].flv",
			"LogLevel": 5,
			"SnapRestSeconds": 15,
			"SplitRestSeconds": 60,
			"CommanderPoolSize": 1,
			"ParserMonitorRestSeconds": 300,
			"DouyinCookie": "__ac_nonce=06245c89100e7ab2dd536; __ac_signature=_02B4Z6wo00f01LjBMSAAAIDBwA.aJ.c4z1C44TWAAEx696;",
			"KuaishouCookie": "did=web_d86297aa2f579589b8abc2594b0ea985"
		},
		"message": ""
	}`

	var tmp interface{}
	jsoniter.UnmarshalFromString(s, &tmp)
	// log.Println(tmp)

	var c Config
	bs, _ := jsoniter.Marshal(tmp)
	jsoniter.Unmarshal(bs, &c)

	log.Println(c)

}

type Config struct {
	// portal
	PortalUsername string `conf:"default:olive"`
	PortalPassword string `conf:"default:olive"`

	// core
	SaveDir                  string `conf:"default:/"`
	OutTmpl                  string `conf:"default:[{{ .StreamerName }}][{{ .RoomName }}][{{ now | date \"2006-01-02 15-04-05\"}}].flv"`
	LogLevel                 uint32 `conf:"default:5"`
	SnapRestSeconds          uint   `conf:"default:15"`
	SplitRestSeconds         uint   `conf:"default:60"`
	CommanderPoolSize        uint   `conf:"default:1"`
	ParserMonitorRestSeconds uint   `conf:"default:300"`

	// tv
	DouyinCookie   string `conf:"default:__ac_nonce=06245c89100e7ab2dd536; __ac_signature=_02B4Z6wo00f01LjBMSAAAIDBwA.aJ.c4z1C44TWAAEx696;"`
	KuaishouCookie string `conf:"default:did=web_d86297aa2f579589b8abc2594b0ea985"`
}
