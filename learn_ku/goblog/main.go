// 每一段 Go 程序都 必须 属于一个包。而 main 包在 Go 程序中有特殊的位置
// 一个标准的可执行的 Go 程序必须有 package main 的声明。
package main

// import 关键词引入 Go Package
import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
}

func main() {
	// 用以指定处理 HTTP 请求的函数
	http.HandleFunc("/", handlerFunc)
	// 用以监听本地 3000 端口并提供服务
	http.ListenAndServe(":3000", nil)
}