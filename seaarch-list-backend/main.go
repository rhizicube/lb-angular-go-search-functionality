package main

import (
	"search-list/controllers"
	"search-list/models"

	"github.com/gin-gonic/gin"
)

func main() {
	sr := gin.Default() // initialize new Default router within sr variable

	models.ConnectDataBase()

	sr.GET("/api/movies", controllers.FindMovies)
	sr.POST("api/movie", controllers.CreateMovie)
	sr.DELETE("api/user/:id", controllers.DeleteMovie)
	sr.POST("/api/sendMovies", controllers.SendMoviesHandler)
	sr.Run("0.0.0.0:8000") // run our server
}
