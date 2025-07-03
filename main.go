package main

import (
	"contestScrap/scrap"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	return r
}

func main() {
	r := setupRouter()
	r.GET("/", func(c *gin.Context){
		contests, err := scrap.GetContests()
		for _, contest := range(contests){
			fmt.Println(contest)
		}
		if(err != nil){
			c.JSON(400, err)
		}
		c.JSON(200, contests)
	})
	r.Run("0.0.0.0:8080")
}
