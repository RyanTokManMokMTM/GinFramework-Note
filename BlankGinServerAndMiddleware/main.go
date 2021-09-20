package main

import "github.com/gin-gonic/gin"

func main() {
	//use gin.New() -> to create a blank new gin server without any middleware
	server := gin.New()

	//add middleware ourselves
	server.Use(gin.Logger()) //at a logger(default) as middleware to server
	server.Use(gin.Recovery()) //at a Recovery(default) as middleware to server

	server.GET("")
	//grouping the route
	server.Run(":8080")
}
