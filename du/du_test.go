/**
 * File: du_test
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/8/26 12:01 下午
 **/

package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestWailDir(t *testing.T)  {
	for i:=0; i<8;i++{
		s :=rand.Int31n(6)
		fmt.Println(s)
	}
}
