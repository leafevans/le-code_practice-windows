package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("Name: %v\tAge: %v\tScore: %v",
		s.Name, s.Age, s.Score)
}

func (s *Student) SetInfo(name string, age, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("打印方法")
}

func PrintStructField(s any) {
	// 判断参数是否为结构体类型
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if (t.Kind() != reflect.Struct) &&
		(t.Elem().Kind() != reflect.Struct) {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	field0 := t.Elem().Field(0)
	fmt.Printf("%v\n", field0)
	fmt.Println(field0.Name)
	fmt.Println(field0.Type)
	fmt.Println(field0.Tag)
	fmt.Println(field0.Index)
	fmt.Println(field0.Offset)

	fmt.Println("------------------------------------")

	field1, ok := t.Elem().FieldByName("Age")
	if ok {
		fmt.Printf("%v\n", field1)
		fmt.Println(field1.Name)
		fmt.Println(field1.Type)
		fmt.Println(field1.Tag)
		fmt.Println(field1.Index)
		fmt.Println(field1.Offset)
	}

	fmt.Println("------------------------------------")

	fmt.Println(t.Elem().NumField())

	fmt.Println("------------------------------------")

	fmt.Println(v.Elem().FieldByName("Name"))
	fmt.Println(v.Elem().FieldByName("Age"))
	fmt.Println(v.Elem().NumField())

	fmt.Println("------------------------------------")

	for i := range t.Elem().NumField() {
		fmt.Printf("属性名称：%v 属性值：%v 属性类型：%v 属性 Tag：%v\n",
			t.Elem().Field(i).Name, v.Elem().Field(i),
			t.Elem().Field(i).Type, t.Elem().Field(i).Tag.Get("json"))
	}
}

func PrintStructFn(s any) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if (t.Kind() != reflect.Struct) &&
		(t.Elem().Kind() != reflect.Struct) {
		fmt.Println("传入的参数不是一个结构体")
	}

	// 按照字母顺序来排列。
	method0 := t.Method(0)
	fmt.Println(method0.Name)
	fmt.Println(method0.Type)

	method1, ok := t.MethodByName("GetInfo")

	if ok {
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}
	fmt.Println("------------------------------------")

	v.Method(1).Call(nil)
	fmt.Println("------------------------------------")

	v.MethodByName("Print").Call(nil)
	fmt.Println("------------------------------------")

	fmt.Println(v.MethodByName("GetInfo").Call(nil))
	fmt.Println("------------------------------------")

	params := []reflect.Value{
		reflect.ValueOf("平泽唯"),
		reflect.ValueOf(23),
		reflect.ValueOf(99),
	}
	v.MethodByName("SetInfo").Call(params)
	fmt.Println(v.MethodByName("GetInfo").Call(nil))
}

func main() {
	stu := Student{
		"平泽唯",
		16,
		100,
	}
	// PrintStructField(&stu)
	PrintStructFn(&stu)
}
