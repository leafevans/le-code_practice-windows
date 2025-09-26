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

// type Computer struct {
// 	Name string
// }

// func (c Computer) Work(usb Usber) {
// 	usb.Start()
// 	usb.Stop()
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

//		computer := Computer{}
//		computer.Work(phone)
//	}
// package main

// import "fmt"

// func show(a any) {
// 	fmt.Printf("Type: %T, Value: %v\n", a, a)
// }

// func main() {
// 	var x any
// 	s := "Hello World!"
// 	x = s
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)

// 	i := 100
// 	x = i
// 	fmt.Printf("Type: %T, Value: %v\n", x, x)

//		b := true
//		x = b
//		fmt.Printf("Type: %T, Value: %v\n", x, x)
//	}
// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// type Student struct {
// 	Id     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// 	Sno    string `json:"sno"`
// }

//	func main() {
//		var studentInfo = make(map[string]any)
//		studentInfo["name"] = "平泽唯"
//		studentInfo["age"] = 18
//		studentInfo["married"] = false
//		fmt.Println(studentInfo)
//		s := &Student{
//			Id:     1,
//			Gender: "女",
//			Name:   "平泽唯",
//			Sno:    "s0001",
//		}
//		fmt.Printf("%#v\n", s)
//		jsonByte, _ := json.Marshal(s)
//		fmt.Println(string(jsonByte))
//		str := `{"Id":1,"Gender":"女","Name":"平泽唯","Sno":"s0001"}`
//		s2 := Student{}
//		err := json.Unmarshal([]byte(str), &s2)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Printf("%#v\n%v\n", s2, s2.Name)
//	}
// package main

// import "fmt"

//	func main() {
//		var slice = []any{"平泽唯", 20, true, 32.2}
//		fmt.Println(slice...)
//	}
// package main

// import "fmt"

//	func main() {
//		var x any = "Hello World!"
//		v, ok := x.(string)
//		if ok {
//			fmt.Println(v)
//		} else {
//			fmt.Println("类型断言失败！")
//		}
//	}
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
// 		fmt.Println("类型断言失败")
// 	}
// }

//	func main() {
//		justifyType(10)
//		justifyType("Hello World!")
//	}
// package main

// import "fmt"

// type Usber interface {
// 	Start()
// 	Stop()
// }

// type Phone struct {
// 	Name string
// }

// func (p *Phone) Start() {
// 	fmt.Println(p.Name, "开机")
// }

// func (p *Phone) Stop() {
// 	fmt.Println(p.Name, "关机")
// }

// func main() {
// 	// 错误写法
// 	// phone1 := Phone{
// 	// 	Name: "小米手机",
// 	// }
// 	// var p1 Usber = phone1
// 	// p1.Start()

//		// 正确写法
//		var p2 Usber = &Phone{
//			Name: "苹果手机",
//		}
//		p2.Start()
//	}
// package main

// import "fmt"

// type Animaler interface {
// 	SetName(string)
// }

// type Peter interface {
// 	GetName() string
// }

// type Dog struct {
// 	Name string
// }

// func (d *Dog) SetName(name string) {
// 	d.Name = name
// }

// func (d Dog) GetName() string {
// 	return d.Name
// }

//	func main() {
//		// 初始化一个结构体。
//		dog := &Dog{
//			Name: "Peter",
//		}
//		var d1 Animaler = dog
//		var d2 Peter = dog
//		d1.SetName("哈基米")
//		fmt.Println(d2.GetName())
//	}
// package main

// type A interface {
// 	SetName(string)
// }

// type B interface {
// 	GetName() string
// }

//	type C interface {
//		A
//		B
//	}
// package main

// import "fmt"

// type Animaler interface {
// 	SetName(string)
// }

// type Peter interface {
// 	GetName() string
// }

// type Dog struct {
// 	Name string
// }

// func (d *Dog) SetName(name string) {
// 	d.Name = name
// }

// func (d Dog) GetName() string {
// 	return d.Name
// }

// func main() {
// 	dog := &Dog{Name: "布鲁斯"}

// 	var d1 Animaler = dog
// 	var d2 Peter = dog

//		fmt.Println(d2.GetName())
//		d1.SetName("布莱恩")
//		fmt.Println(d2.GetName())
//	}
// package main

// import "fmt"

// type Sayer interface {
// 	say()
// }

// type Mover interface {
// 	move()
// }

// type Animaler interface {
// 	Sayer
// 	Mover
// }

// type Cat struct {
// 	name string
// }

// func (c Cat) say() {
// 	fmt.Println("基米多")
// }

// func (c Cat) move() {
// 	fmt.Println("南北绿")
// }

//	func main() {
//		var a Animaler = Cat{
//			name: "花花",
//		}
//		a.move()
//		a.say()
//	}
// package main

// import "fmt"

// type Runner interface {
// 	Run()
// }

// type Swimer interface {
// 	Swim()
// }

// type Animal interface {
// 	Runner
// 	Swimer
// }

// type Cat struct {
// 	Name string
// }

// func (c Cat) Run() {
// 	fmt.Println("基米跑步")
// }

// func (c Cat) Swim() {
// 	fmt.Println("基米游泳")
// }

//	func main() {
//		var cat Animal = &Cat{Name: "哈基米"}
//		cat.Run()
//		cat.Swim()
//	}
package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

func main() {
	userinfo := make(map[string]any)
	userinfo["username"] = "平泽唯"
	userinfo["age"] = 18
	userinfo["hobby"] = []string{"偷闲", "吃甜品"}

	fmt.Println(userinfo["hobby"].([]string)[1])

	address := Address{
		Name:  "田井中律",
		Phone: 1234567890,
	}
	userinfo["address"] = address
	a, ok := userinfo["address"].(Address)
	if ok {
		fmt.Println(a.Name)
	}
}
