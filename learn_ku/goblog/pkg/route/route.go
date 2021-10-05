package route

import (
	"github.com/gorilla/mux"
)

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	var router *mux.Router
	url, err := router.Get(routeName).URL(pairs...)
	if err != nil {
		// logger.LogError(err)
		return ""
	}

	return url.String()
}