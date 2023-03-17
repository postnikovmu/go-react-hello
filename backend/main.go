package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string
}

func hello(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	message := Message{Text: "Hello from Go!"}
	res.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(message)
	res.Write(jsonData)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil) //locally
}
