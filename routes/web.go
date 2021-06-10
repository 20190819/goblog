package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yangliang4488/goblog/app/http/controllers"
	"github.com/yangliang4488/goblog/app/http/middlewares"
)

var pc *controllers.PageController = new(controllers.PageController)
var ac *controllers.ArticlesController = new(controllers.ArticlesController)

// 用户认证
var auth *controllers.AuthController = new(controllers.AuthController)

func RegisterWebRoutes(router *mux.Router) {

	router.Use(middlewares.ForceHtml)
	router.Use(middlewares.StartSession)

	// 静态页面
	router.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	router.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	router.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	// 文章
	router.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/{id:[1-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("GET").Name("articles.delete")

	router.HandleFunc("/auth/register", auth.Register).Methods("GET").Name("auth.register")
	router.HandleFunc("/auth/register", auth.DoRegister).Methods("POST").Name("auth.doRegister")
	// 登录
	router.HandleFunc("/auth/login", auth.Login).Methods("GET").Name("auth.login")
	router.HandleFunc("/auth/login", auth.DoLogin).Methods("POST").Name("auth.doLogin")

}
