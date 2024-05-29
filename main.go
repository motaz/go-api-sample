package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/addorder", AddOrder)
	http.HandleFunc("/removeorder", RemoveOrder)
	http.HandleFunc("/getorders", GetOrders)
	http.HandleFunc("/", About)
	port := "10022"
	fmt.Println("http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error while listening: ", err.Error())
	}
}
