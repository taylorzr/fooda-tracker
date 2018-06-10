package main

import (
	"github.com/taylorzr/fooda-tracker/http"
)

func main() {
	http.Router().Run(":8080")
}
