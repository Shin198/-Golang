package main

import "net/http"

//call poiter here to handler
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responeWithJSON(w, 200, struct{}{})
}
