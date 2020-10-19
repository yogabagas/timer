package main

import (
	"fmt"
	"time"
)

var messages = make(chan string)

func main() {
	go stop()
	timer(messages)

	fmt.Println()
}

func timer(req chan string) {

	ticker := time.NewTicker(time.Second)
	// var ti time.Time

	for {
		select {
		case <-req:
			return
		case t := <-ticker.C:
			h, m, s := t.Clock()
			fmt.Println("TIME", fmt.Sprintf("%d:%d:%d", h, m, s))
		}
	}

}

func stop() {
	// reader := bufio.NewReader(os.Stdin)
	// msg, _ := reader.ReadString('\n')
	var msg string
	fmt.Scanln(&msg)

	messages <- msg
}
