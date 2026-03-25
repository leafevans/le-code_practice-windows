// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // 处理器函数（HandlerFunc）
// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello 哈基米！")
// }

// func main() {
// 	// 注册路由
// 	http.HandleFunc("/", sayHello)

//		// 启动服务，监听端口
//		err := http.ListenAndServe(":9090", nil)
//		if err != nil {
//			fmt.Println("启动服务失败:", err)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello 哈基米!")
// }

//	func main() {
//		http.HandleFunc("/", sayHello)
//		err := http.ListenAndServe(":9090", nil)
//		if err != nil {
//			fmt.Println("启动服务失败:", err)
//		}
//	}
package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	query := r.URL.Query()
	name := query.Get("name")

	r.ParseForm()
	age := r.PostForm.Get("age")

	body, _ := io.ReadAll(r.Body)
	var jsonData map[string]any
	json.Unmarshal(body, &jsonData)

	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]any{
		"status": "ok",
		"name":   name,
		"age":    age,
		"json":   jsonData,
	})
	w.Write(resp)
}
