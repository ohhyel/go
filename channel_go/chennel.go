package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan string)
	go writePing(c)
	go readPing(c)

	var input string
	fmt.Scanln(&input)
}

func readPing(strings chan string) {
	fmt.Println("readPing")
	for {
		msg := <-strings
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func writePing(chanstrings chan string) {
	for i := 0; ; i++  {
		fmt.Printf("writePing %d\n", i)
		chanstrings <- "ping"
	}
}


