/**
 * File: test1
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2021/10/28 8:19 下午
 **/

package slice

import (
	"sort"
	"unsafe"
)

var a = []float64{45,1,6,3,79,8,4,6,15}

func sortFloat642Fast(a []float64)  {

	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	var c (*[1 << 20]int)
	//c := [5]int{}
	//d := c[:1:2] //从切片建立数组 [from:len:cap] from开始位置 len 切片长度 cap 切片容量

	dd := unsafe.Pointer(&a[1])
	sort.Ints(b)
}
