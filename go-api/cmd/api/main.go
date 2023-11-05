package main

import (
	"fmt"
	"net/http"

	"github.com/Shin198/go-api/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handlers(r)

	fmt.Println("Starting API services...")

	err := http.ListenAndServe("localhost:8088", r)
	if err != nil {
		log.Error(err)
	}
}
