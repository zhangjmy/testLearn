/**
 * File: main
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/9/22 8:34 下午
 **/

package main

import (
"fmt"
"net"
)

const MaxBufferSize = 128    //单次读取最大长度
const HeaderLen = 5    //包头长度

func main() {
	tcpAddr := ":15000"
	listener, err := net.Listen("tcp4", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	//启动客户端程序
	go send()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

//处理从通道接收的数据
func accept(acceptData chan string) {
	for {
		value, isOk := <- acceptData
		if !isOk {
			break
		}
		parse([]byte(value))
	}
}

func parse(data []byte) {
	//将四字节的包头转为string
	head := string(data[:4])
	//将一字节的包体长度转为十进制整形
	size := int(data[4])
	//将包体转为string
	body := string(data[5:])
	fmt.Println(fmt.Sprintf("接收的数据：包头为：%s，包体大小%d字节，值为：%s", head, size, body))
}

//处理tcp请求
func handle(conn net.Conn) {
	defer conn.Close()
	//创建接收数据的通道
	acceptData := make(chan string, 10)
	defer close(acceptData)
	go accept(acceptData)

	reader := Reader{
		Conn:  conn,
		Buff:  make([]byte, MaxBufferSize),
		Start: 0,
		End:   0,
		BuffLen: MaxBufferSize,
		HeaderLen: HeaderLen,
	}
	err := reader.Read(acceptData)
	if err != nil {
		fmt.Println("read data error:", err)
	}
}
