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

	for {
		select {
		case <-req:
			return
		case t := <-ticker.C:
			fmt.Println("TIME", fmt.Sprintf("%d:%d:%d", t.Hour(), t.Minute(), t.Second()))
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
