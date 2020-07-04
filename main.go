package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Static("/StaticData", "./TgBuild/StaticData")
	router.LoadHTMLGlob("./TgBuild/index.html")
	router.GET("/", loadMainPage)
	log.Panic(router.Run(fmt.Sprintf("%v", os.Getenv("PORT"))))
}

func loadMainPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{})
}