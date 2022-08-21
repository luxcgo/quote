package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Masterminds/sprig"
)

// var tmpl = template.Must(template.New("filename3").Funcs(GetFuncMap()).
// Parse(`{{ .Live.GetPlatformCNName }}/{{ .HostName | filenameFilter }}/[{{ now | date "2006-01-02 15-04-05"}}][{{ .HostName | filenameFilter }}][{{ .RoomName | filenameFilter }}].flv`))

var tmpl1 = template.Must(template.New("filename").Funcs(sprig.TxtFuncMap()).
	Parse(`[{{ .StreamerName }}][{{ .RoomName }}][{{ now | date "2006-01-02 15-04-05"}}].flv`))

func main1() {
	buf := new(bytes.Buffer)
	if err := tmpl1.Execute(buf, nil); err != nil {
		panic(fmt.Sprintf("failed to render filename, err: %v", err))
	}
	println(buf.String())
}

func main() {
	buf := new(bytes.Buffer)
	info := &struct {
		StreamerName string
		RoomName     string
	}{
		StreamerName: "StreamerName",
		RoomName:     "room",
	}
	// info := map[string]string{
	// 	"HostName": "host",
	// }
	if err := tmpl1.Execute(buf, info); err != nil {
		panic(fmt.Sprintf("failed to render filename, err: %v", err))
	}
	if err := tmpl1.Execute(buf, info); err != nil {
		panic(fmt.Sprintf("failed to render filename, err: %v", err))
	}
	println(buf.String())
}

type Info struct {
	Live
	HostName, RoomName          string
	Status, Listening, Recoding bool
}

type Live struct {
}

func (l Live) GetPlatformCNName() (int, error) {
	return 62, nil
}
