package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/task/internal/controller/song"
	"github.com/nurzzaat/task/internal/repository"
	"github.com/nurzzaat/task/pkg"

	_ "github.com/nurzzaat/task/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(app pkg.Application, router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	songController := song.SongController{
		SongRepository: repository.NewSongRepository(app.Sql),
	}

	songRoutes := router.Group("/song")
	{
		songRoutes.POST("", songController.CreateSong)
		songRoutes.GET("", songController.GetAll)
		songRoutes.GET("/:songId", songController.GetByID)
		songRoutes.PATCH("/:songId", songController.UpdateSong)
		songRoutes.DELETE("/:songId", songController.DeleteSong)
	}

}
