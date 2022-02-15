/**
 * File: selectt
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/8/25 8:00 下午
 **/

package goroutine

import (
	"fmt"
	"os"
	"time"
)

func maint() {
	fmt.Println("Commencing countdown.  Press return to abort.")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
				abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("-----quit-----")
		//return
	}
	fmt.Println("-------send fire-----")

	ch := make(chan int,1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}
