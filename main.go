package main

import "log"

func main() {
	s := "123你好"
	for _, c := range s {
		log.Printf("%c\n", uint8(c))
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		log.Printf("%c\n", c)
	}
	//log.Println(len("123你好"))
}
