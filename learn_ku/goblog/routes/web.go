package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"goblog/app/http/middlewares"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	// 静态页面
	pc := new(controllers.PagesController)

	// 用以指定处理 HTTP 请求的函数，/ 意味着任意路径
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")

	// 精确匹配
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 自定义 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)

	// 文章相关页面
	ac := new(controllers.ArticlesController)

	// 正则匹配
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")

	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")

	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete")

	// 用户注册
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", auc.Register).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", auc.DoRegister).Methods("POST").Name("auth.doregister")

	// 用户认证
	r.HandleFunc("/auth/login", auc.Login).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", auc.DoLogin).Methods("POST").Name("auth.dologin")

	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 开始会话
	r.Use(middlewares.StartSession)
}
