package controllers

import (
	"fmt"
	"net/http"

	"github.com/yangliang4488/goblog/app/models/article"
	"github.com/yangliang4488/goblog/pkg/logger"
	"github.com/yangliang4488/goblog/pkg/route"
	"gorm.io/gorm"
)

type ArticlesController struct{}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		fmt.Fprint(w, "读取成功，文章标题 —— "+article.Title)
	}
}
