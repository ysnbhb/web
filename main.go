package main

import (
	"net/http"

	web "web/func" // package creat by team
)

func main() {
	http.HandleFunc("/", web.Print) // if i was in url / use func Print
	http.HandleFunc("/ascii-art", web.Handel_input)
	http.HandleFunc("/Download", web.Download)
	http.ListenAndServe(":9090", nil) // this func for run server
}
