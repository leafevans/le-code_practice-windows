package main

import (
	"fmt"
	"net/http"
)

// 处理器函数（HandlerFunc）
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello 哈基米！")
}

func main() {
	// 注册路由
	http.HandleFunc("/", sayHello)

	// 启动服务，监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("启动服务失败:", err)
	}
}
