package router

import (
	"net/http"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/DevelopNaoki/webapp-user/model"
)

func Router(engine *gin.Engine) {
	p, _ := os.Getwd()
	fmt.Print(p)
	engine.Static("/css", "src/css")
	engine.Static("/js", "src/js")

	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
        engine.GET("/register", func(c *gin.Context) {
                c.HTML(http.StatusOK, "register.html", gin.H{})
        })

	engine.POST("/login", func(c *gin.Context) {
		model.Login()
        })
}
