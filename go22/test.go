package main

import (
	"errors"
	"fmt"
)

func sum(a int, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("input can't be minus")
	}
	sum = a + b
	err = nil
	return sum, err
}

func sum1(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

func main() {
	result, err := sum(-11, 2)
	if err != nil {
		fmt.Println((err))
	} else {
		fmt.Println(result)
	}

	fmt.Println(sum1(1, 2, 3, 4, 5))

	sum2 := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum2(1, 2))

	cl := closure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	age := Age(25)

	age.String()
	age.Modify()
	age.String()

	p := person{
		age:  20,
		name: "AgarthaSF",
		addr: address{
			province: "Anhui",
			city:     "Chaohu",
		},
	}

	fmt.Println(p.age, p.name, p.addr.city)
	printString(&p)
	printString(p.addr)

	p1 := NewPerson("张三")

	var s fmt.Stringer
	s = p1
	p2 := s.(*person)
	fmt.Println(p2)
	a, ok := s.(address)
	if ok {

		fmt.Println(a)

	} else {

		fmt.Println("s不是一个address")

	}
	// if ok {

	// 	fmt.Println(a)

	// } else {

	// 	fmt.Println("s不是一个address")

	// }

}

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}

func (age *Age) Modify() {
	*age = Age(30)
}

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

// type Stringer interface {
// 	String() string
// }

func (p *person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func (addr address) String() string {
	return fmt.Sprintf("the province is %s, city is %s", addr.province, addr.city)
}

func NewPerson(name string) *person {
	return &person{name: name}
}

// 结构体
type errorString struct {
	s string
}

// 返回error接口
func New(text string) error {
	return &errorString{text}
}

// 实现error接口
func (e *errorString) Error() string {
	return e.s
}
