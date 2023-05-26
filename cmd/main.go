package main

import (
	"fmt"
	"net/http"

	"github.com/youssef-182/production-server/pkg/router"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GOT REQUEST ON /")
	w.Write([]byte("WELL HELLO THERE"))
}

func main() {
	r := router.Setup()

	http.ListenAndServe(":1337", r)
}
