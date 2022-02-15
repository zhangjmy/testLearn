/**
 * File: client
 * DESC:
 * Author: zhangjintao
 * Email: zhangjintao01@toyoogame.com
 * DATE:  2021/7/17 6:08 下午
 **/

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}