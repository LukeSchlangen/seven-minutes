package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello Gophers Hello schooltactic watermelon"))
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}