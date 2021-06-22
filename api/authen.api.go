package api

import (
	_ "time"

	"github.com/gin-gonic/gin"
	"main/model"
)

//ถ้าตอนนี้ ยังไม่ได้ใช้ก็ให้ void เอาไว้ก่อน โดยการใส่ _ นำหน้า

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
		// มีการเปลี่ยน Login->login มีคำอธิบายว่า จริง ๆ แล้วเราไม่ได้เรียกจากข้างนอก เพราะฉนั้น เราใช้ตัวเล็กก็ได้
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

// Login - login api
func login(c *gin.Context) {
	//เปลี่ยน 401 -> 200
	c.JSON(200, gin.H{"result": "login"})
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil{
		//ถ้าการ Bind ไม่มี error
		c.JSON(200, gin.H{"result": "register", "data": user})
	}
	
}
