package main

import (
	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	router.POST("/cs", ScoringController)
	router.Run("0.0.0.0:8000")	//Listen adn serve on 0.0.0.0:8000
}
