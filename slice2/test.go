/**
 * File: test
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/11/17 2:49 下午
 **/

package main

import (
	"fmt"
)

func main()  {
	s :=[]int{21,43,3,5,4,3,66,7,88}
	fmt.Println(s)
	s1 := s[0:3]
	fmt.Println(s1)
	s2 := s[2:4]
	fmt.Println(s2)

	s3:= s[2:2]
	fmt.Println(s3)
}
