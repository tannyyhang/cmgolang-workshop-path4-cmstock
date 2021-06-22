package api

import (
	"net/http"

	"github.com/gin-gonic/gin")


func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		// ตรงนี้เคยมี error เกิดขึ้นเนื่องจาก
		transactionAPI.GET("/transaction", getTransaction) //ตรงนี้เคยใช้ POST เลยเปลืี่ยนเป็น GET แทน (จะสังเกตุเห็นว่า่ ใช้ /product เหมือนกัน )
		transactionAPI.POST("/transaction", createTransaction) //ดังนั้น จะต้อง พิจารณา  GET กับ POST ให้ดี
	}
}

// Login - login api
func getTransaction(c *gin.Context) {
	c.String(http.StatusOK,"List Transaction")
}

func createTransaction(c *gin.Context) {
	c.String(http.StatusOK,"Create Transaction")
}
