package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Static("/StaticData", "./TgBuild/StaticData")
	router.LoadHTMLGlob("./TgBuild/index.html")
	router.GET("/", loadMainPage)
	log.Panic(router.Run(":3301"))
}

func loadMainPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}