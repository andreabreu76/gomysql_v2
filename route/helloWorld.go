package route

import (
	"github.com/andreabreu76/gomysql_v2/handler"
	"github.com/gorilla/mux"
)

func SetupRoutes() {

	router := mux.NewRouter()

	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/", handler.Hello).Methods("GET")

}
