package main

import (
	"fmt"
	"net/http"
)

// 路由方法
func router(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(w)
	fmt.Println(r)
}

func main() {
	// 静态路由解析
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/image/", http.FileServer(http.Dir("public")))
	// 动态路由解析
	http.HandleFunc("/", router)
	// 创建web服务器
	http.ListenAndServe(":8000", nil)
}
