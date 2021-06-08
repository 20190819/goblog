package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/yangliang4488/goblog/app/http/middlewares"
	"github.com/yangliang4488/goblog/bootstrap"
	"github.com/yangliang4488/goblog/pkg/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 欢迎来到 my goblog！</h1>")
		fmt.Fprint(w, time.Now().String())
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

type Article struct {
	Title, Body string
	ID          int64
}
type ArticleFormatData struct {
	Title, Body string
	URL         *url.URL
	Errors      error
}

var db *sql.DB

var router *mux.Router

func getRouteVariable(key string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[key]
}

func main() {
	// 初始化数据库
	db = database.DB
	database.Initialize()

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	router.Use(middlewares.ForceHtml)
	homeUrl, _ := router.Get("home").URL()
	fmt.Println("homeUrl>>>", homeUrl)

	articleUrl, _ := router.Get("articles.show").URL("id", "123")
	fmt.Println("articleUrl>>>", articleUrl)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
