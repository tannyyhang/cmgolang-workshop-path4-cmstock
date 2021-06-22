package api

import (
	"time"
	"main/db"
	"main/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	c.JSON(200, gin.H{"result": "login"})
	var user model.User

	if c.ShouldBind(&user) == nil {
		var queryUser model.User		
		if err := db.GetDB().First(&queryUser, "username = ?", user.Username).Error; err != nil {							
			c.JSON(200, gin.H{"result": "nok", "error": err})
		}else if (checkPasswordHash(user.Password, queryUser.Password) == false){
			c.JSON(200, gin.H{"result": "nok", "error": "invalid password"})
		}else{			
			c.JSON(200, gin.H{"result": "ok", "data": user})
		}

	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}	
}


func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}


//func register(c *gin.Context) {
	//var user model.User
	//if c.ShouldBind(&user) == nil{
		//ถ้าการ Bind ไม่มี error
		//c.JSON(200, gin.H{"result": "register", "data": user})
	//}
	
//}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil{
		//ถ้าการ Bind ไม่มี error
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		if err := db.GetDB().Create(&user).Error; err != nil{
			c.JSON(200, gin.H{"result": "nok", "error": err})
		}else{
			c.JSON(200, gin.H{"result": "ok", "data": user})
		}
		
	}
	
}