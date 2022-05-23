package utils

import (
	"fmt"
	"net/http"
)

func JsonResponse(msg interface{}, code int, w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintf(w, `{"code":%d,"message":"%s", "url_path":"%s"}`, code, msg, r.URL.Path)
}
