package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting...")
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method, req.URL.Path)
		res.WriteHeader(200)
		res.Write([]byte(`{"ok":true}`))
		res.Header().Add("content-type", "application/json")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		panic(err)
	}
}
