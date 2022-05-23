package route

import (
	"net/http"

	"github.com/andreabreu76/gomysql_v2/handler"
)

func SetupRoutes() {

	http.HandleFunc("/", handler.Hello)

}
