package main

import (
	"razorpay/controllers"
	"razorpay/initializer"

	"github.com/gin-gonic/gin"
)

func init(){
	initializer.LoadEnv()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/app", controllers.ExecuteRazorpay)
	router.GET("/payment-Status", controllers.PaymentStatus)


	router.Run("localhost:8080")

}
