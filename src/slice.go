package main

import (
	"fmt"
)

func slice() {
	fmt.Println("--------------------slice--------------------")

	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:5]

	fmt.Printf("value:%v, len:%v, cap:%v\n", t, len(t), cap(t))

	src := []int{1, 2, 3, 4, 5}
	dst := src[:0]
	fmt.Printf("value:%v, len:%v, cap:%v\n", src, len(src), cap(src))
	fmt.Printf("value:%v, len:%v, cap:%v\n", dst, len(dst), cap(dst))
	for _, v := range src {
		if v%2 == 0 {
			dst = append(dst, v)
		}
	}
	fmt.Printf("value:%v, len:%v, cap:%v\n", src, len(src), cap(src))
	fmt.Printf("value:%v, len:%v, cap:%v\n", dst, len(dst), cap(dst))

	for i := len(dst); i < len(src); i++ {
		src[i] = 0
	}
	fmt.Printf("value:%v, len:%v, cap:%v\n", src, len(src), cap(src))
	fmt.Printf("value:%v, len:%v, cap:%v\n", dst, len(dst), cap(dst))
}
