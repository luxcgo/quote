package main

import (
	"sync"
)

var pool *sync.Pool

type Info struct {
	Name string
}
type Person struct {
	Direct string
	*Info
}

func (this *Person) clone() *Person {
	// tmp := *this
	// return &tmp
	return this
}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			return &Person{
				Info: &Info{},
			}
		},
	}
}

func main() {
	initPool()
	p := pool.Get().(*Person)
	p.Direct = "1direct"
	p.Name = "first"
	pool.Put(p)
	p2 := pool.Get().(*Person)
	p2.Name = "second"
	p2.Direct = "2direct"
	// pool.Put(p2)
	println(p.Name)
	println(p.Direct)
	println(p2.Name)
	println(p2.Direct)
}
