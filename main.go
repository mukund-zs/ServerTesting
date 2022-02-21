package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}
	name := req.URL.Query().Get("name")
	if name == "" {
		log.Println("Name is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := w.Write([]byte("My name: " + name))
	if err != nil {
		log.Println(err)

		return
	}
	w.WriteHeader(http.StatusOK)
}

//func main() {
//	http.HandleFunc("/hello", helloHandler)
//	log.Fatal(http.ListenAndServe(":8000", nil))
//}
