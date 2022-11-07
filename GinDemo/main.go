package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jc0803kevin/go-study/GinDemo/controller"
	"net/http"
)

func main() {
	router := gin.Default()
	// 静态资源加载，本例为css,js以及资源图片
	//StaticFS 是加载一个完整的目录资源
	router.StaticFS("/public", http.Dir("E:\\workspace\\gowork\\src\\github.com\\jc0803kevin\\go-study\\GinDemo\\web\\static"))
	//StaticFile 是加载单个文件
	router.StaticFile("/logo.jpg", "./web/static/kaiyuan10nian.jpg")
	// 导入所有模板，多级目录结构需要这样写
	router.LoadHTMLGlob("E:\\workspace\\gowork\\src\\github.com\\jc0803kevin\\go-study\\GinDemo\\web\\templete\\*")
	// website分组
	v := router.Group("/")
	{
		v.GET("/index.html", controller.IndexController)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}