package main

import (
	"bufio"
	"fmt"
	"os"
)

type message struct {
	body      string
	messageId int
}

type hub struct {
	reciever chan message
	sender   chan message
}

func (m *hub) readMeaage() {
	for {
		select {
		case r, ok := <-m.reciever:
			fmt.Println("受信メッセージ", r)
			if ok {
				m.sender <- message{body: "ok", messageId: r.messageId}
			}
		}

	}
}

func main() {
	h := &hub{reciever: make(chan message), sender: make(chan message)}
	go h.readMeaage()

	// 標準入力
	sc := bufio.NewScanner(os.Stdin)
	var id = 0
	for {
		sc.Scan()
		inputText := sc.Text()
		switch inputText {
		case "exit":
			close(h.reciever)
			close(h.sender)
		default:
			h.reciever <- message{body: inputText, messageId: id}
			id += 1
		}
		select {
		case r := <-h.sender:
			if r.body == "ok" {
				fmt.Println("メッセージを受け付けました")
			}
		}
	}
}
