package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type person struct {
	Name string `json:"name" bson:"b_name"`
	Age  uint   `json:"age" bson:"b_age"`
}

func (p person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
	p := person{Name: "AgarthaSF", Age: 20}

	// struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}

	respJSON := "{\"Name\":\"张三\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)

	pt := reflect.TypeOf(p)
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fmt.Println("realize fmt.Stringer? ", pt.Implements((stringerType)))
	fmt.Println("realize io.Writer? ", pt.Implements((writerType)))

	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n", sf.Name, sf.Tag.Get("json"))
		fmt.Printf("字段%s上,bson tag为%s\n", sf.Name, sf.Tag.Get("bson"))
	}

	fmt.Println(struct2json(p))

	pa := person{Name: "AgarthaSF", Age: 20}
	pva := reflect.ValueOf(pa)
	mPrint := pva.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("Login")}
	mPrint.Call(args)

}

func (p person) Print(prefix string) {
	fmt.Printf("%s: Name is %s, Age is %d\n", prefix, p.Name, p.Age)
}

func struct2json(p person) string {
	// p := person{Name: "AgarthaSF", Age: 20}
	pv := reflect.ValueOf(p)
	pt := reflect.TypeOf(p)

	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")
	num := pt.NumField()
	for i := 0; i < num; i++ {
		jsonTag := pt.Field(i).Tag.Get("json")
		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")

		// 获取字段的值
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))

		if i < num-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	return jsonBuilder.String()
}
