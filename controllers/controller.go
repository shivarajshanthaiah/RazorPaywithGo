package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)

type PageVariables struct {
	AppointmentID string
	Name          string
	Email         string
	Amount        int
}

func ExecuteRazorpay(c *gin.Context) {
	page := &PageVariables{}

	page.Name = "Shivaraj Shanthaiah"
	page.Email = "shivaraj3652@gmail.com"
	page.Amount = 350000
	client := razorpay.NewClient(os.Getenv("key_id"), os.Getenv("key_secret"))

	data := map[string]interface{}{
		"amount":   page.Amount,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	fmt.Println("", body)

	if err != nil {
		fmt.Println("payment not initiated")
		os.Exit(1)
	}

	value := body["id"]
	str := value.(string)
	HomePageVar := PageVariables{
		AppointmentID: str,
		Name:          page.Name,
		Email:         page.Email,
		Amount:        page.Amount/100,
	}
	c.HTML(http.StatusOK, "app.html", HomePageVar)
}

func PaymentStatus(c *gin.Context) {
	paymentid:=c.Query("paymentid")
	appointmentid := c.Query("appointmentid")
	name := c.Query("name")

	fmt.Println("paymentid", paymentid)
	fmt.Println("appointmentid", appointmentid)
	fmt.Println("name", name)

}
