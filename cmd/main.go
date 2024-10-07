package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/task/internal/controller"
	"github.com/nurzzaat/task/pkg"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.

//	@host		localhost:1232
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	log.SetFormatter(&ecslogrus.Formatter{})

	path := "./logs/song.log"

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	defer app.Sql.Close()

	ginRouter := gin.Default()
	ginRouter.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS ,HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	controller.Setup(app, ginRouter)

	ginRouter.Run(fmt.Sprintf(":%s", app.Env.PORT))
}
