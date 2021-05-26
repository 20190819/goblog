package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/yangliang4488/goblog/pkg/database"
	"github.com/yangliang4488/goblog/pkg/logger"
	"github.com/yangliang4488/goblog/pkg/route"
	"github.com/yangliang4488/goblog/pkg/types"

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

func articlesShowhandler(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouteVariable("id", r)
	article, err := getArticleById(id)
	if err != nil {
		if err == sql.ErrNoRows {
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

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT * from articles"
	rows, err := db.Query(sql)
	logger.LogError(err)

	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Body)
		logger.LogError(err)
		articles = append(articles, article)
	}

	err = rows.Err()
	logger.LogError(err)
	// 3. 加载模板
	tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
	logger.LogError(err)

	// 4. 渲染模板，将所有文章的数据传输进去
	tmpl.Execute(w, articles)
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "请提供正确的数据！")
	} else {
		title := r.PostForm.Get("title")
		body := r.PostForm.Get("body")
		lastInsertID, err := saveArticleDB(title, body)
		if lastInsertID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+types.Int64ToString(lastInsertID))
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500服务器内部错误")
		}
	}

}

func articlesCreateHandler(rw http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<title>创建文章 —— 我的技术博客</title>
		</head>
		<body>
			<form action="/articles?test=data" method="post">
				<p><input type="text" name="title"></p>
				<p><textarea name="body" cols="30" rows="10"></textarea></p>
				<p><button type="submit">提交</button></p>
			</form>
		</body>
	</html>
	`
	storeUrl, _ := router.Get("articles.store").URL()
	fmt.Fprint(rw, html, storeUrl)
}

func articlesEditHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 获取ID
	id := route.GetRouteVariable("id", r)
	// 2. 读取文章
	article, err := getArticleById(id)

	if err == nil {
		updateUrl, _ := router.Get("articles.update").URL("id", id)
		data := ArticleFormatData{
			Title:  article.Title,
			Body:   article.Body,
			URL:    updateUrl,
			Errors: nil,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)
		tmpl.Execute(w, data)
	} else {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章未找到~")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器错误")
		}

	}
}

func articlesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 获取ID
	id := route.GetRouteVariable("id", r)
	// 2. 读取文章
	_, err := getArticleById(id)
	if err == nil {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		validateErr := validateArticleFormData(title, body)
		if len(validateErr) == 0 {
			// 通过验证
			query := "UPDATE articles SET title = ?, body = ? WHERE id = ?"
			rs, err := db.Exec(query, title, body, id)
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "服务器错误")
				return
			}

			if n, _ := rs.RowsAffected(); n > 0 {

				showURL, _ := router.Get("articles.show").URL("id", id)
				http.Redirect(w, r, showURL.String(), http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		} else {
			fmt.Fprint(w, "表单验证未通过")
		}
	} else {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章未找到~")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器错误")
		}
	}

}

func articlesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := getArticleById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "文章未找到~")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器错误")
		}
	} else {
		rowAffected, err := article.delete()
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "服务器错误")
		} else if rowAffected > 0 {
			fmt.Fprint(w, "删除成功")
		}
	}

}

func (article Article) delete() (RowsAffected int64, err error) {
	rs, err := db.Exec("DELETE FROM articles WHERE id = " + types.Int64ToString(article.ID))
	if err != nil {
		return 0, err
	} else {
		if n, _ := rs.RowsAffected(); n > 0 {
			return n, nil
		} else {
			return 0, nil
		}
	}
}

func getArticleById(id interface{}) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err
}

func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)
	// 1. 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	// 2.验证内容
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}
	return errors
}

func forceHtmlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html;charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(rw, r)
	})
}

var db *sql.DB

func saveArticleDB(title string, body string) (int64, error) {
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)

	stmt, err = db.Prepare("INSERT INTO articles (title, body) VALUES(?,?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	rs, err = stmt.Exec(title, body)

	if id, err = rs.LastInsertId(); id > 0 {
		return id, nil
	} else {
		return 0, err
	}

}

var router *mux.Router

func main() {
	// 初始化数据库
	db = database.DB
	database.Initialize()
	// 初始化路由
	route.Initialize()
	router = route.Router

	router.HandleFunc("/", defaultHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowhandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")

	router.HandleFunc("/articles/{id:[1-9]+}/edit", articlesEditHandler).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesUpdateHandler).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).Methods("GET").Name("articles.delete")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.Use(forceHtmlMiddleware)

	homeUrl, _ := router.Get("home").URL()
	fmt.Println("homeUrl>>>", homeUrl)

	articleUrl, _ := router.Get("articles.show").URL("id", "123")
	fmt.Println("articleUrl>>>", articleUrl)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
