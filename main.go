package main

import (
	"fmt"
	"time"
	"github.com/taylorzr/fooda-tracker/db"
	"github.com/taylorzr/fooda-tracker/http"
	"github.com/taylorzr/fooda-tracker/fooda"
	"github.com/taylorzr/fooda-tracker/hipchat"
)

func main() {
	go func() {
		for {
			notify()
			// time.Sleep(5 * time.Minute)
			time.Sleep(5 * time.Second)
		}
	}()

	http.Router().Run(":8081")
}

// You could move this to the fooda package if you injected the dependencies db and hipchat
// We would need to probably create structs with methods for db and hipchat, maybe make them
// interfaces
// Actually, could probably move the eternal loop into fooda root package as well
func notify() {
	orders := db.TodaysOrders()

	users := db.AllUsers()

	forgetters := fooda.Forgetters(users, orders)

	if len(forgetters) > 0 {
		hipchat.Message("test", fmt.Sprintf("%v", forgetters))
	} else {
		fmt.Println("No forgetters")
	}
}
