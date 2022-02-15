/**
 * File: test1.go
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/9/24 11:05 上午
 **/

package main

import (
	"fmt"
)

func main() {
	//maxInt := int(^uint(0) >> 1)
	//fmt.Printf("-------%d", maxInt)
	TestTx()
}

func TestTx() {
	var s1 uint32 = 100
	var s2 uint32 = 200
	s3 := s1 - s2
	fmt.Printf("s3:%v", s3)
	s := "3" > "1"
	fmt.Printf("%v", s)
}
