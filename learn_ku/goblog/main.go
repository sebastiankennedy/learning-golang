// 每一段 Go 程序都 必须 属于一个包。而 main 包在 Go 程序中有特殊的位置
// 一个标准的可执行的 Go 程序必须有 package main 的声明。
package main

// import 关键词引入 Go Package
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var router = mux.NewRouter()

// 使用中间件
func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置表头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("X-App-Name", "GoBlog")
		// 继续处理
		next.ServeHTTP(w, r)
	})
}

// 首页控制器
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog！</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	// URL 路径参数解析为键值对应的 Map
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章 ID："+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// 解析错误，这里应该有错误处理
		fmt.Fprint(w, "请提供正确的数据！")
		return
	}

	title := r.PostForm.Get("title")

	fmt.Fprintf(w, "POST PostForm: %v <br>", r.PostForm)
	fmt.Fprintf(w, "POST Form: %v <br>", r.Form)
	fmt.Fprintf(w, "title 的值为: %v", title)
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			request.URL.Path = strings.TrimSuffix(request.URL.Path, "/")
		}

		next.ServeHTTP(writer, request)
	})
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>创建文章 —— 我的技术博客</title>
</head>
<body>
    <form action="%s?test=data" method="post">
        <p><input type="text" name="title"></p>
        <p><textarea name="body" cols="30" rows="10"></textarea></p>
        <p><button type="submit">提交</button></p>
    </form>
</body>
</html>
`
	storeURL, _ := router.Get("articles.store").URL()
	fmt.Fprintf(w, html, storeURL)
}

func main() {
	// 用以指定处理 HTTP 请求的函数，/ 意味着任意路径
	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")

	// 精确匹配
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	// 正则匹配
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")

	// 指定不同的 HTTP 方法、路由名称
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 通过命名路由获取 URL 示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleURL: ", articleURL)

	// 注册中间件
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
