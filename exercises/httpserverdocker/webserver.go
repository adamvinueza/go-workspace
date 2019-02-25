package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandlerHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Result struct {
	Method string
}

func indexHandlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(Result{r.Method})
	io.WriteString(w, string(result))
}
