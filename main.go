package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var AppVersion = "v2.4"

//var Pr = true
var builtBy = "unknown"

func init() {
	fmt.Printf("Powered by go-olive/olive %s by %s\n", AppVersion, builtBy)
	version := flag.Bool("v", false, "print olive version")
	flag.Parse()
	if *version {
		fmt.Println(version)
		os.Exit(0)
	}
}

func main() {
	fmt.Println("haha")
}

var flagPath = flag.String("path", "test", "path to traverse and rename.")

func renamePath(path string, f os.FileInfo, err error) (e error) {
	if strings.HasPrefix(f.Name(), "renameMe") {
		dir := filepath.Dir(path)
		base := filepath.Base(path)
		newname := filepath.Join(dir, strings.Replace(base, "renameMe", "newName", 1))
		log.Printf("mv %q %q\n", path, newname)
		os.Rename(path, newname)
	}
	return
}

func init() {
	flag.Parse()
}

func main1() {
	dir := "/Users/lucas/github/go-acwing/problems"
	err := filepath.WalkDir(dir, func(fpath string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}
		base := filepath.Base(fpath)
		s := strings.Split(base, "-")[0]
		n := filepath.Join(filepath.Dir(fpath), s)
		// fmt.Println(n)
		os.Rename(fpath, n)
		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}

}
