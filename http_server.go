package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
	url := r.RequestURI
	fmt.Println(url)
	w.Write([]byte(url+"\n"))
}

func main() {
	http.HandleFunc("/",HelloHandler)
	http.ListenAndServe("0.0.0.0:2003",nil)
}
