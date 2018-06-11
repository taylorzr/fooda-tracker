package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/taylorzr/fooda-tracker/http"
)

func main() {
	go func() {
		for {
			notify()
			sleep(5 * time.Minute)
		}
	}()

	http.Router().Run(":8081")
}

// You could move this to the fooda package if you injected the dependencies 'db' and 'hipchat'
func notify() {
	orders := db.TodaysOrders()

	users := db.AllUsers()

	forgetters := fooda.Forgetters(users, orders)

	hipchat.Notify(forgetters)
}
