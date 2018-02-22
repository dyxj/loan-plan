package main

import (
	"github.com/dyxj/loan-plan/api"
	"log"
	"net/http"
	"time"
)

func main() {
	rt := http.NewServeMux()
	api.InitAPIRoutes(rt)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        rt,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
