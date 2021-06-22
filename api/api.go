package api

import(
	"github.com/gin-gonic/gin"
	//_ "main/db"
)

func Setup(router *gin.Engine){
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)
}