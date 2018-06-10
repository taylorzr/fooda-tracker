package http

import(
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/taylorzr/fooda-tracker/fooda"
	"github.com/taylorzr/fooda-tracker/db"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/", hello)
	router.GET("/orders", getOrder)
	router.POST("/orders", createOrder)
	return router
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"response":  "Hello World",
	})
}

func getOrder(c *gin.Context) {
	c.JSON(200, db.GetOrder("zach.taylor@avant.com"))
}

func createOrder(c *gin.Context) {
	var order fooda.Order

	err := c.BindJSON(&order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.CreateOrder(&order)

	c.JSON(201, order)
}
