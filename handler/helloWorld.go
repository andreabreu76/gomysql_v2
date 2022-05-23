package handler

import (
	"net/http"

	"github.com/andreabreu76/gomysql_v2/utils"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse("Hello World", http.StatusOK, w, r)
}
