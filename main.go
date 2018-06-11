package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/taylorzr/fooda-tracker/http"
	"time"
)

func main() {
	go func() {
		// TODO: Store last time we run this in database
		for {
			if time.Now().Hour() == 9 && time.Now().Minute() > 45 {
				fmt.Println("Hello world")
			} else {
				gin.Logger().Info("Not time...")
			}

			time.Sleep(1 * time.Second)
		}
	}()

	http.Router().Run(":8081")
}
