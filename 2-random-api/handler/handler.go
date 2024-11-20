package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type RandomHandler struct{}

func NewRandomHandler(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.Handle("/random", handler.random())
}

func (h *RandomHandler) random() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		w.Write([]byte(fmt.Sprintf("%d", rand.Intn(6) + 1)))
	}
}
