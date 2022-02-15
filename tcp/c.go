/**
 * File: c
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/9/22 8:37 下午
 **/

package main
import (
	"fmt"
	"net"
	"time"
)

func send() {
	addr := ":15000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	head := []byte("head")
	bodies := []string{
		"I am not afraid of tomorrow for I have seen yesterday and I love today.",
		"If you want to understand today, you have to search yesterday.",
		"You never know what you can do till you try.",
		"A good name keeps its luster in the dark.",
	}
	for {
		var messages []byte
		for _, b := range bodies{
			body := []byte(b)
			oneMessage := append(head, byte(len(body)))
			oneMessage = append(oneMessage, body...)
			messages = append(messages, oneMessage...)
		}
		_, err := conn.Write(messages)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(5*time.Second)
	}
}
