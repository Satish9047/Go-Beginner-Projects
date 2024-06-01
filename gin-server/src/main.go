package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("Hello World!")

	//create a router
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.String(200, "Hello World!")
	})

	router.GET("/hello", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.Run(":8080")
	
}
