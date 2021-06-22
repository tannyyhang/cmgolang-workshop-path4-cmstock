package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//เนื่องจากเราต้องการให้ func ตัวนี้ สามารถเรียกใช้จากข้างนอกได้ เราต้องขึ้นต้นด้วยตัวใหญ่ -> General....
//GeneralInterceptor1  - call this method to add interceptor
func GeneralInterceptor1(c *gin.Context){
	token := c.Query("token")
	if(token == "1234"){
		c.Next() // ถ้า token ตรง c.Next() ก็คือการบอกว่า ไปต่อได้
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}
}