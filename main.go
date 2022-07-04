package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	println("hi")
	println("squash")
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
