package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
) 
type Game struct { 

	Time 	string 	`json:"id"`
}
func init_game (c *gin.Context) { 
	var new_game Game
	if err := c.BindJSON(&new_game) ; err != nil { 
		return 
	} 
	
	
}
func main () { 
	router := gin.Default()
	router.GET("/init_game" , )
}