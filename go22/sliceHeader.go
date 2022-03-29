package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// ss := []string{"AgarthaSF", "张三"}
	// fmt.Println("length: ", len(ss), " capacity: ", cap(ss))
	// ss = append(ss, "李四")
	// fmt.Println(ss)
	// fmt.Println("length: ", len(ss), " capacity: ", cap(ss))

	a1 := [2]string{"AgarthaSF", "Suifeng"}
	s1 := a1[0:1]
	s2 := a1[:]
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	s := "AgarthaSF"
	fmt.Printf("S's memory address: %d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b's memory address: %d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("s3's memory address: %d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s3)).Data)
	fmt.Println(s, string(b), s3)

	// 通过指针使用byte的内存转换string，没有额外开销
	s4 := *(*string)(unsafe.Pointer(&b))
	fmt.Printf("s4's memory address: %d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s4)).Data)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b1 := *(*[]byte)(unsafe.Pointer(sh))

	fmt.Println(sh, string(b1))

}
