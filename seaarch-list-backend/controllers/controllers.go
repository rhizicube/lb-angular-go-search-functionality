package controllers

import (
	"net/http"
	"search-list/models"
	"time"

	"github.com/gin-gonic/gin"
)

//Get list of all users
func FindMovies(ctx *gin.Context) {
	time.Sleep(2 * time.Second)
	var users []models.Movie
	models.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// Create User using Create()
func CreateMovie(ctx *gin.Context) {
	//validate User
	var input models.CreateMovie
	if err := ctx.ShouldBindJSON(&input); err != nil { // syntax??
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//create user
	user := models.Movie{Name: input.Name}
	models.DB.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

//delete the user
func DeleteMovie(ctx *gin.Context) {
	//find the user
	var user models.Movie
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//delete the user from database using delete() provided by middleware
	models.DB.Delete(&user)
	//return deeted user
	ctx.JSON(http.StatusOK, gin.H{"data": true}) //nopoint in returning deleted object
}

func SendMoviesHandler(ctx *gin.Context) {
	time.Sleep(2 * time.Second)
	var input models.CreateMovie
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var movies []models.Movie
	// smallName := strings.ToLower(input.Name)
	models.DB.Where("name = ?", input.Name).Find(&movies)
	ctx.JSON(http.StatusOK, gin.H{"data": movies})
}
