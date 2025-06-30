package main

import (
	"contestScrap/scrap"
	"fmt"
	"github.com/gin-gonic/gin"
)


func setupRouter() *gin.Engine {

	r := gin.Default()

	
	return r
}

func main() {
	r := setupRouter()
	contests, err := scrap.GetContests()
	for _, contest := range(contests){
		fmt.Println(contest)
	}
	fmt.Println(err)
	// for _, contest := range contests {
	// 	fmt.Println(contest)
	// }
	// print(err)

	r.Run(":8080")
}
