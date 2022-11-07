package main

import "net/http"

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("success"))
	})

	// 定义需要在 服务ListenAndServe 启动之间
	http.Handle("/hello", &HelloHandler{})
	http.ListenAndServe(":18080", nil)

}

type HelloHandler struct {
}

// 需要定义一个结构体 实现这个 ServeHTTP 这个接口，
func (HelloHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello"))
}
