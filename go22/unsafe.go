package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	Name string
	Age  uint
}

func test() {
	p := new(person)
	pName := (*string)(unsafe.Pointer(p))
	*pName = "AgarthaSF"

	// 进行内存偏移
	pAge := (*uint)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)
}

func main() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)
	test()

	fmt.Println(unsafe.Sizeof(true))

	fmt.Println(unsafe.Sizeof(int8(0)))

	fmt.Println(unsafe.Sizeof(int16(10)))

	fmt.Println(unsafe.Sizeof(int32(10000000)))

	fmt.Println(unsafe.Sizeof(int64(10000000000000)))

	fmt.Println(unsafe.Sizeof(int(10000000000000000)))

	fmt.Println(unsafe.Sizeof(string("AgarthaSF")))

	fmt.Println(unsafe.Sizeof([]string{"AgarthaSF", "张三"}))
}
