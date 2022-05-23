package main

import (
	"net/http"

	"github.com/andreabreu76/gomysql_v2/route"
)

func main() {

	route.SetupRoutes()

	http.ListenAndServe(":8080", nil)
}
