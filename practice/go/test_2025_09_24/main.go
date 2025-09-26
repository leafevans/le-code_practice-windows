// package main

// import "fmt"

// type myInt = int
// type newInt int

// type person struct {
// 	name, citt string
// 	age        int8
// }

// func main() {
// 	var a newInt
// 	var b myInt
// 	fmt.Printf("Type of a: %T\n", a)
// 	fmt.Printf("Type of b: %T\n", b)
// }

// package main

// import "fmt"

// type person struct {
// 	name, city string
// 	age        int
// }

// // func main() {
// // 	var p1 person
// // 	p1.name = "哈基米"
// // 	p1.city = "北京"
// // 	p1.age = 18
// // 	fmt.Printf("p1 = %v\n", p1)
// // 	fmt.Printf("p1 = %#v\n", p1)
// // }

// func main() {
// 	var p2 = new(person)
// 	p2.name = "曼波"
// 	p2.age = 19
// 	p2.city = "上海"

// 	fmt.Printf("%T\n", p2)
// 	fmt.Printf("p2 = %#v\n", p2)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	city string
// 	age  int
// }

// func main() {
// 	p3 := &person{}
// 	fmt.Printf("%T\n", p3)
// 	fmt.Printf("p3 = %#v\n", p3)

// 	p3.name = "哈基米"
// 	p3.age = 30
// 	p3.city = "杭州"
// 	(*p3).age = 40
// 	fmt.Printf("p3 = %#v\n", p3)
// }

// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) PrintInfo() {
// 	fmt.Printf("Name: %v, Age: %v\n", p.name, p.age)
// }

// func main() {
// 	p1 := Person{
// 		"曼波",
// 		25,
// 	}
// 	p1.PrintInfo()
// }

// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) printInfo() {
// 	fmt.Printf("Name: %v\tAge: %v\n", p.name, p.age)
// }

// func (p *Person) setInfo(name string, age int) {
// 	p.name = name
// 	p.age = age
// }

// func main() {
// 	p1 := &Person{
// 		name: "哈牛魔",
// 		age:  29,
// 	}
// 	p1.printInfo()
// 	p1.setInfo("ccb", 20)
// 	p1.printInfo()
// }

// package main

// import "fmt"

// type myInt int

// func (m myInt) SayHello() {
// 	fmt.Println("Hello Int!")
// }

// func main() {
// 	var m1 myInt
// 	m1.SayHello()
// 	m1 = 100
// 	fmt.Printf("%#v\t%T\n", m1, m1)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	fmt.Printf("Current time: %v\n", now)

// 	year := now.Year()
// 	month := now.Month()
// 	day := now.Day()
// 	hour := now.Hour()
// 	minute := now.Minute()
// 	second := now.Second()

// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
// 		year, month, day,
// 		hour, minute, second)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()

// 	fmt.Println(now.Format("2006-01-02 15:04:05"))

// 	fmt.Println(now.Format("2006-01-02 03:04:05"))

// 	fmt.Println(now.Format("2006/01/02 15:04"))
// 	fmt.Println(now.Format("15:04 2006/01/02"))
// 	fmt.Println(now.Format("2006/01/02"))
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	unixTime := now.Unix()
// 	unixNano := now.UnixNano()

// 	fmt.Printf("Current Unix time: %v\n", unixTime)
// 	fmt.Printf("Current Unix nano: %v\n", unixNano)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func unixToTime(timestamp int64) {
// 	timeObj := time.Unix(timestamp, 0)
// 	str := timeObj.Format("2006-01-02 15:04:05")
// 	fmt.Println(str)
// }

// func main() {
// 	unixToTime(1758348998)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func unixToTime(timestamp int64) {
// 	timeObj := time.Unix(timestamp, 0)
// 	str := timeObj.Format("2006-01-02 15:04:05")
// 	fmt.Println(str)
// }

// func main() {
// 	unixToTime(1758348998)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	str := "2025-09-24 15:02:03"
// 	timeTemple := "2006-01-02 15:04:05"
// 	timestamp, _ := time.ParseInLocation(timeTemple, str, time.Local)
// 	fmt.Println(timestamp.Unix())
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	now := time.Now()
// 	later := now.Add(time.Second*10 + time.Hour*2)
// 	fmt.Println(later.Format("2006-01-02 15:04:05"))
// 	fmt.Println(now.Sub(later))
// 	fmt.Println(now.Before(later))
// 	fmt.Println(later.After(now))
// }

// 'package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	n := 0

// 	for v := range ticker.C {
// 		fmt.Println(v)
// 		n++
// 		if n > 5 {
// 			ticker.Stop()
// 			return
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	n := 0
// 	for {
// 		time.Sleep(time.Second)
// 		n++
// 		fmt.Println("test...")
// 		if n > 5 {
// 			return
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// type Student struct {
// 	ID     int
// 	Gender string
// 	Name   string
// 	Sno    string
// }

// func main() {
// 	s1 := Student{
// 		ID:     12,
// 		Gender: "女",
// 		Name:   "平泽唯",
// 		Sno:    "s0001",
// 	}
// 	fmt.Printf("%#v\n", s1)
// 	jsonByte, _ := json.Marshal(s1)
// 	jsonStr := string(jsonByte)
// 	fmt.Println(jsonStr)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	ID     int
// 	Gender string
// 	name   string
// 	Sno    string
// }

// func main() {
// 	str := `{"ID":12, "Gender":"女", "Name":"平泽唯", "Sno":"s00001"}`
// 	var s1 Student
// 	err := json.Unmarshal([]byte(str), &s1)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf("%#v\n", s1)
// 	fmt.Println(s1.name)
// }

// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// // 学生
// type Student struct {
// 	ID     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// }

// // 班级
// type Class struct {
// 	Title    string    `json:"title"`
// 	Students []Student `json:"students"`
// }

// func main() {
// 	// var s = Student{
// 	// 	12,
// 	// 	"女",
// 	// 	"平泽唯",
// 	// 	"s0001",
// 	// }
// 	// fmt.Printf("%#v\n", s)

// 	// jsonByte, _ := json.Marshal(s)
// 	// jsonStr := string(jsonByte)
// 	// fmt.Println(jsonStr)
// 	// c := &Class{
// 	// 	Title:    "0302",
// 	// 	Students: make([]Student, 0, 200),
// 	// }
// 	// for i := range 10 {
// 	// 	s := Student{
// 	// 		ID:     i,
// 	// 		Gender: "男",
// 	// 		Name:   fmt.Sprintf("stu_%v", i),
// 	// 	}
// 	// 	c.Students = append(c.Students, s)
// 	// }
// 	// // fmt.Println(c)
// 	// jsonByte, err := json.Marshal(c)

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// fmt.Println(string(jsonByte))
// 	str := `
// {
//     "title": "0302",
//     "students": [
//         {
//             "id": 0,
//             "gender": "男",
//             "name": "stu_0"
//         },
//         {
//             "id": 1,
//             "gender": "男",
//             "name": "stu_1"
//         },
//         {
//             "id": 2,
//             "gender": "男",
//             "name": "stu_2"
//         },
//         {
//             "id": 3,
//             "gender": "男",
//             "name": "stu_3"
//         },
//         {
//             "id": 4,
//             "gender": "男",
//             "name": "stu_4"
//         },
//         {
//             "id": 5,
//             "gender": "男",
//             "name": "stu_5"
//         },
//         {
//             "id": 6,
//             "gender": "男",
//             "name": "stu_6"
//         },
//         {
//             "id": 7,
//             "gender": "男",
//             "name": "stu_7"
//         },
//         {
//             "id": 8,
//             "gender": "男",
//             "name": "stu_8"
//         },
//         {
//             "id": 9,
//             "gender": "男",
//             "name": "stu_9"
//         }
//     ]
// }
// `
// 	c := &Class{}

// 	err := json.Unmarshal([]byte(str), c)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Printf("%#v\n", c)
// }

// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// type Student struct {
// 	ID     int
// 	Gender string
// 	name   string
// 	Sno    string
// }

// func main() {
// 	s := &Student{
// 		ID:     1,
// 		Gender: "女",
// 		name:   "平泽唯",
// 		Sno:    "s0001",
// 	}
// 	fmt.Printf("%#v\n", s)

// 	jsonByte, _ := json.Marshal(s)
// 	jsonStr := string(jsonByte)
// 	fmt.Println(jsonStr)
// }

// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// 	Sno    string
// }

// func main() {
// 	jsonStr := `{"Id":1, "Gender":"女", "Name":"平泽唯", "Sno":"s0001"}`
// 	s := &Student{}

// 	err := json.Unmarshal([]byte(jsonStr), s)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Printf("%#v\t%v\n", s, s.Name)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Student struct {
// 	Id     int    `json:"id"`
// 	Gender string `json:"gender"`
// 	Name   string `json:"name"`
// 	Sno    string `json:"sno"`
// }

// func main() {
// 	s := &Student{
// 		Id:     1,
// 		Gender: "女",
// 		Name:   "平泽唯",
// 		Sno:    "s0001",
// 	}
// 	fmt.Printf("%#v\n", s)

// 	jsonByte, _ := json.Marshal(s)
// 	jsonStr := string(jsonByte)
// 	fmt.Println(jsonStr)
// }

// package main

// import (
// 	"fmt"

// 	json "github.com/json-iterator/go"
// )

// // 学生
// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// }

// // 班级
// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	c := &Class{
// 		Title:    "0302",
// 		Students: make([]Student, 0, 200),
// 	}

// 	for i := range 10 {
// 		s := Student{
// 			Name:   fmt.Sprintf("stu_%02d", i),
// 			Gender: "女",
// 			Id:     i,
// 		}
// 		c.Students = append(c.Students, s)
// 	}

// 	data, err := json.Marshal(c)

// 	if err != nil {
// 		fmt.Println("JSON marshal failed.")
// 		return
// 	}

// 	fmt.Printf("JSON: %s\n", data)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // 学生
// type Student struct {
// 	Id     int
// 	Gender string
// 	Name   string
// }

// // Class
// type Class struct {
// 	Title    string
// 	Students []Student
// }

// func main() {
// 	str := `
// {
//     "Title": "0302",
//     "Students": [
//         {
//             "Id": 0,
//             "Gender": "女",
//             "Name": "stu_00"
//         },
//         {
//             "Id": 1,
//             "Gender": "女",
//             "Name": "stu_01"
//         },
//         {
//             "Id": 2,
//             "Gender": "女",
//             "Name": "stu_02"
//         },
//         {
//             "Id": 3,
//             "Gender": "女",
//             "Name": "stu_03"
//         },
//         {
//             "Id": 4,
//             "Gender": "女",
//             "Name": "stu_04"
//         },
//         {
//             "Id": 5,
//             "Gender": "女",
//             "Name": "stu_05"
//         },
//         {
//             "Id": 6,
//             "Gender": "女",
//             "Name": "stu_06"
//         },
//         {
//             "Id": 7,
//             "Gender": "女",
//             "Name": "stu_07"
//         },
//         {
//             "Id": 8,
//             "Gender": "女",
//             "Name": "stu_08"
//         },
//         {
//             "Id": 9,
//             "Gender": "女",
//             "Name": "stu_09"
//         }
//     ]
// }
// `

// 	c := &Class{}

// 	err := json.Unmarshal([]byte(str), c)

// 	if err != nil {
// 		fmt.Println("JSON unmarshal failed!")
// 		return
// 	}

// 	fmt.Printf("%#v\n", c)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func makeSuffix(suffix string) func(string) string {
// 	return func(name string) string {
// 		if !strings.HasSuffix(name, suffix) {
// 			return name + suffix
// 		}
// 		return name
// 	}
// }

// func main() {
// 	jpgFn := makeSuffix(".jpg")
// 	txtFn := makeSuffix(".txt")
// 	fmt.Println(jpgFn("test"))
// 	fmt.Println(txtFn("test"))
// }

package main

import (
	"fmt"
	"test_2025_09_24/calc"
	"test_2025_09_24/tools"
)

func init() {
	fmt.Println("Hello World!")
}

func main() {
	x := 10.2
	y := 20.3
	fmt.Println(calc.Add(x, y))
	fmt.Println(tools.Mul(x, y))
}
