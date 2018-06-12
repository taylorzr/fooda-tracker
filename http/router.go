package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/taylorzr/fooda-tracker/db"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/", hello)
	router.GET("/orders/:email", getOrder)
	router.POST("/orders", createOrder)
	return router
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "Hello World",
	})
}

func getOrder(c *gin.Context) {
	c.JSON(200, db.GetOrder(c.Param("email")))
}

func createOrder(c *gin.Context) {
	var orderRequest OrderRequest

	err := c.BindJSON(&orderRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	order := db.CreateOrder(orderRequest.Email, orderRequest.OrderedAt)

	c.JSON(201, order)
}

type (
	OrderRequest struct {
		OrderedAt time.Time `json:"ordered_at"`
		Email string `json:"email"`
	}
)
