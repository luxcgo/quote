package main

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/coocood/freecache"
	l "github.com/luxcgo/lifesaver/log"
)

func main() {
	cacheSize := 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	key := []byte("1")
	val := []byte("def")
	expire := 10 // expire in 60 seconds
	cache.Set(key, val, expire)
	cache.Set([]byte("2"), val, expire)
	fmt.Println(cache.EntryCount())

	time.Sleep(time.Second)
	// got, err := cache.Get(key)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", got)
	// }
	iter := cache.NewIterator()
	for entry := iter.Next(); entry != nil; entry = iter.Next() {
		l.Logger.Info(string(entry.Key))
		l.Logger.Info(string(entry.Value))
	}

	fmt.Println(cache.EntryCount())
}
