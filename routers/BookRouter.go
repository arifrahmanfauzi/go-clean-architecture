package routers

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/controllers"
	"go-clean-architecture/repository"
	"go-clean-architecture/usecase"
	"gorm.io/gorm"
)

func BookRouter(group *gin.RouterGroup, db *gorm.DB) {
	createBookRepository := repository.ProviderBookRepository(db)
	createBookUsecase := usecase.NewBookUsecase(createBookRepository)
	BookHandler := controllers.NewBookController(createBookUsecase)
	group.GET("/books", BookHandler.Fetch)
	group.POST("/book", BookHandler.Store)
	group.PUT("/book/:id", BookHandler.Update)
}
