/**
 * File: json
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/12/27 10:47 AM
 **/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := 20 / 100
	fmt.Printf("%v", s)
	var s1 string
	v(&s1)

	fmt.Printf("ret:%v", s1)
}

func v(s interface{}) {
	old := "dddd"
	ss := reflect.TypeOf(s)
	sk := ss.String()
	//skk := ss.Kind()
	//ssv := reflect.ValueOf(s)
	if ss.String() == "*string" {
		s1 := s.(*string)
		*s1 = old
		//sv := reflect.ValueOf(s)
		//ssv.Bytes()
		return
	}
	fmt.Printf("%v\n", sk)
}
