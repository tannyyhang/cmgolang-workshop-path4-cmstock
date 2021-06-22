package api

import (
	"net/http"
	"main/interceptor"
	"github.com/gin-gonic/gin")


	// โดยปรกติ เราจะไม่ทำ func myInterceptor ไว้ในที่เดียวกับตัวที่ handler
func myInterceptor(c *gin.Context){
	token := c.Query("token")
	if(token == "1234"){
		c.Next() // ถ้า token ตรง c.Next() ก็คือการบอกว่า ไปต่อได้
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}
}

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		// ตรงนี้เคยมี error เกิดขึ้นเนื่องจาก
		productAPI.GET("/product",interceptor.GeneralInterceptor1, getProduct) //ตรงนี้เคยใช้ POST เลยเปลืี่ยนเป็น GET แทน (จะสังเกตุเห็นว่า่ ใช้ /product เหมือนกัน )
		productAPI.POST("/product", createProduct) //ดังนั้น จะต้อง พิจารณา  GET กับ POST ให้ดี
	}
}

// Login - login api
func getProduct(c *gin.Context) {
	//เปลี่ยน 401 -> 200
	c.JSON(200, gin.H{"result": "get procudt"})
}

func createProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "create product"})
}
