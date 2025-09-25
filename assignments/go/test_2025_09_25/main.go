// package main

// import "fmt"

// var x int8 = 10

// const pi = 3.14

// func init() {
// 	fmt.Println(x)
// }

// func main() {
// 	fmt.Println("Hello World!")
// }

// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// type Student struct {
// 	Id   uint32 `json:"id"`
// 	Name string `json:"name"`
// 	Age  uint8  `json:"age"`
// 	Sex  string `json:"sex"`
// }

// func main() {
// 	s := Student{
// 		1,
// 		"平泽唯",
// 		18,
// 		"女",
// 	}
// 	data, err := json.Marshal(s)
// 	if err != nil {
// 		fmt.Println("Error!")
// 	}
// 	fmt.Println(string(data))
// }

// package main

// import "fmt"

// type Usber interface {
// 	start()
// 	stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p Phone) start() {
// 	fmt.Println(p.Name, "启动")
// }

// func (p Phone) stop() {
// 	fmt.Println(p.Name, "关机")
// }

// func (p Phone) run() {
// 	fmt.Println(p.Name, "运行")
// }

// func main() {
// 	var p Usber = Phone{
// 		Name: "iPhone",
// 	}
// 	p.start()
// 	if phone, ok := p.(Phone); ok {
// 		phone.run()
// 	}
// 	p.stop()
// }

// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Computer struct {
// }

// func (c Computer) Work(usb Usber) {
// 	usb.Start()
// 	usb.Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p Phone) Start() {
// 	fmt.Println(p.Name, "启动")
// }

// func (p Phone) Stop() {
// 	fmt.Println(p.Name, "关机")
// }

// func main() {
// 	c := Computer{}
// 	p := Phone{
// 		Name: "OPPO",
// 	}
// 	c.Work(p)
// }

// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p Phone) Start() {
// 	fmt.Println(p.Name, "开机")
// }

// func (p Phone) Stop() {
// 	fmt.Println(p.Name, "关机")
// }

// type Camera struct {
// }

// func (c Camera) Start() {
// 	fmt.Println("相机 开机")
// }

// func (c Camera) Stop() {
// 	fmt.Println("相机 关机")
// }

// func main() {
// 	var phone Usber = Phone{
// 		Name: "一加手机",
// 	}
// 	phone.Start()
// 	phone.Stop()

// 	camera := Camera{}
// 	var c Usber = camera
// 	c.Start()
// 	c.Stop()
// }

// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p Phone) Start() {
// 	fmt.Println(p.Name, "开机")
// }

// func (p Phone) Stop() {
// 	fmt.Println(p.Name, "关机")
// }

// type Computer struct {
// 	Name string
// }

// func (c Computer) Work(usb Usber) {
// 	usb.Start()
// 	usb.Stop()
// }

// func main() {
// 	phone := Phone{
// 		Name: "小米手机",
// 	}
// 	computer := Computer{}
// 	computer.Work(phone)
// }

// package main

// import "fmt"

// type A interface{}

// func main() {
// 	var a A
// 	var str = "Hello Golang!"
// 	a = str // 表示字符串实现了该接口。
// 	fmt.Printf("value = %v, type = %T\n", a, a)

// 	var num = 20
// 	a = num // 表示整型实现了该接口。
// 	fmt.Printf("value = %v, type = %T\n", a, a)

// 	var flag = true
// 	a = flag // 表示布尔型实现了该接口。
// 	fmt.Printf("value = %v, type = %T\n", a, a)
// }

// package main

// import "fmt"

// func show(a any) {
// 	fmt.Printf("value = %v, type = %T\n", a, a)
// }

// func main() {
// 	var a any
// 	var str = "Hello Golang!"
// 	a = str // 表示字符串实现了该接口。
// 	show(a)

// 	var num = 20
// 	a = num // 表示整型实现了该接口。
// 	show(a)

// 	var flag = true
// 	a = flag // 表示布尔型实现了该接口。
// 	show(a)
// }

// package main

// import "fmt"

// func main() {
// 	m1 := make(map[string]any)
// 	m1["username"] = "张三"
// 	m1["age"] = 20
// 	m1["married"] = true
// 	fmt.Println(m1)
// 	s1 := []any{1, "2", true}
// 	fmt.Println(s1)
// 	fmt.Println(s1[0].(string))
// }

// package main

// import "fmt"

// func main() {
// 	var a any
// 	a = 11
// 	value, ok := a.(string)
// 	if ok {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("断言失败！")
// 	}
// }

// package main

// import "fmt"

// func justifyType(x any) {
// 	switch x.(type) {
// 	case string:
// 		fmt.Println("string 类型")
// 	case int:
// 		fmt.Println("int 类型")
// 	case bool:
// 		fmt.Println("bool 类型")
// 	default:
// 		fmt.Println("传入错误……")
// 	}
// }

// func main() {
// 	justifyType(10)
// 	justifyType("Hello World!")
// }

// package main

// import "fmt"

// func main() {
// 	var x any
// 	s := "Hello World!"
// 	x = s
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)

// 	i := 100
// 	x = i
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)

// 	b := true
// 	x = b
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)
// }

// package main

// import "fmt"

// func main() {
// 	var x any = "Hello Golang!"
// 	v, ok := x.(string)
// 	if ok {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println("类型断言失败！")
// 	}
// }

package main

import (
	"fmt"
)

func justifyType(x any) {
	switch x.(type) {
	case string:
		fmt.Println("string 类型")
	case int:
		fmt.Println("int 类型")
	case bool:
		fmt.Println("bool 类型")
	case float32:
		fmt.Println("float32 类型")
	case float64:
		fmt.Println("float64 类型")
	default:
		fmt.Println("Unsupport type!")
	}
}

func main() {
	justifyType(10)
	justifyType("Hello World")
}
