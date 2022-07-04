package main

import "log"

func main() {
	topoOrder := make([]int, 5)
	change(&topoOrder)
	log.Println(topoOrder)
}

func change(s *[]int) {
	*s = append(*s, 1, 2, 3)
	(*s)[0] = -1
}
