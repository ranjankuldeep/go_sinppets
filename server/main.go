package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleFunc)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("handler started")
	defer log.Printf("handler ended")

	ctx := r.Context()

	select {
	case <-ctx.Done():
		log.Printf(ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	case <-time.After(5 * time.Second):
		fmt.Fprint(w, "hello there from the server")
	}
}
