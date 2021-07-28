// 每一段 Go 程序都 必须 属于一个包。而 main 包在 Go 程序中有特殊的位置
// 一个标准的可执行的 Go 程序必须有 package main 的声明。
package main

// import 关键词引入 Go Package
import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>你好，这里是 GoBlog。</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
			"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+"<p>如有疑惑，请联系我们。</p>")
	}
}

func main() {
	// 用以指定处理 HTTP 请求的函数，/ 意味着任意路径
	http.HandleFunc("/", handlerFunc)
	// 用以监听本地 3000 端口并提供服务
	http.ListenAndServe(":3000", nil)
}
