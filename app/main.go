package main

import (
	"github.com/gin-gonic/gin"

	"github.com/DevelopNaoki/webapp-user/router"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("src/templates/*.html")

	router.Router(engine)

	engine.Run(":8080")
}
