package main

import "log"

func main() {
	a := 1
	b := 2
	a, b = b, a
	log.Println(a, b)
}
