package api

import "github.com/gin-gonic/gin"

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		// ตรงนี้เคยมี error เกิดขึ้นเนื่องจาก
		productAPI.GET("/product", getProduct) //ตรงนี้เคยใช้ POST เลยเปลืี่ยนเป็น GET แทน (จะสังเกตุเห็นว่า่ ใช้ /product เหมือนกัน )
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
