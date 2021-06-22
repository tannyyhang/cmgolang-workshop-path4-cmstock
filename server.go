package main

import(
	//"main/api"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.Static("/images","./uploaded/images")
	router.Run(":8081")
}
