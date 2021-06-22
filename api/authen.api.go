package api

import "github.com/gin-gonic/gin"

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		// มีการเปลี่ยน Login->login มีคำอธิบายว่า จริง ๆ แล้วเราไม่ได้เรียกจากข้างนอก เพราะฉนั้น เราใช้ตัวเล็กก็ได้
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

// Login - login api
func login(c *gin.Context){
	//เปลี่ยน 401 -> 200
	c.JSON(200, gin.H{"result": "login"})
}

func register(c *gin.Context){
	c.JSON(200, gin.H{"result": "register"})
}