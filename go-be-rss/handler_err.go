package main

import "net/http"

//call poiter here to handler
func handlerErr(w http.ResponseWriter, r *http.Request) {
	responeWithJSON(w, 400, "Something wrong")
}
