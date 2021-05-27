package controllers

import (
	"fmt"
	"html/template"
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

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	articles, err := article.GetAll()
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		// 2. 加载模板
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		logger.LogError(err)
		// 3. 渲染模板，将所有文章的数据传输进去
		tmpl.Execute(w, articles)
	}
}
