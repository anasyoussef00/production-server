package main

import (
	"net/http"

	"github.com/youssef-182/production-server/pkg/router"
)

func main() {
	r := router.Setup()

	http.ListenAndServe(":1337", r)
}
