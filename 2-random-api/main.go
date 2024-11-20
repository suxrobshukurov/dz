package main

import (
	"fmt"
	"go/adv/2-random-api/handler"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	handler.NewRandomHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Listening on port 8081")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
