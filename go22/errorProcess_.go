package main

import (
	"errors"
	"fmt"
	"strconv"
)

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

type WalkRun interface {
	Walk()
	Run()
}

func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}

func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}

type commonError struct {
	errorCode int    // 错误码
	errorMsg  string //错误信息
}

func (ce *commonError) Error() string {
	return ce.errorMsg
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{
			errorCode: 1,
			errorMsg:  "a,b不能为负数",
		}
	} else {
		return a + b, nil
	}
}

type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	return e.err.Error() + e.msg
}

func main() {
	p := person{
		name: "AgarthaSF",
		age:  20,
		addr: address{
			province: "Hubei",
			city:     "Wuhan",
		},
	}
	p.Run()

	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	sum, err := add(-1, 2)

	//如果返回的 ok 为 true，说明 error 断言成功
	//说明发生了错误
	if cm, ok := err.(*commonError); ok {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

	e := errors.New("原始错误")
	w := fmt.Errorf("Wrap错误: %w", e)
	fmt.Println(w)

	fmt.Println(errors.Unwrap(w))

	fmt.Println(e == w)
	fmt.Println(errors.Is(w, e))

	var cm *commonError
	if errors.As(err, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}
}
