package routes

import (
	// Import the Gin library
	"github.com/gin-gonic/gin"
)

type CountryController struct{}

// HelloWorldController will hold the methods to the
// Default controller handles returning the hello world JSON response
func (h *CountryController) Default(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello world, climate change is real"})
}

func (h *CountryController) FindOne(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello world, climate change is real"})
}

func (h *CountryController) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello world, climate change is real"})
}
