/**
 * File: file
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/11/24 4:05 下午
 **/

package main

import (
	"fmt"
	"os"
)

func main()  {
	file := "/Users/dominic/work/yoozoo/idlewarserver/draw.txt"
	slog, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil{
		fmt.Println(err)
	}
	_, err = slog.WriteString("ssdfsdfa\n")
	if err != nil{
		fmt.Println(err)
	}
}
