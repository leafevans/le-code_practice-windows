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

package main

import "fmt"

type Usber interface {
	Start()
	Stop()
}

type Computer struct {
}

func (c Computer) Work(usb Usber) {
	usb.Start()
	usb.Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "关机")
}

func main() {
	c := Computer{}
	p := Phone{
		Name: "OPPO",
	}
	c.Work(p)
}
