package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/addorder", AddOrder)
	http.HandleFunc("/removeorder", RemoveOrder)
	http.HandleFunc("/getorders", GetOrders)
	fmt.Println("http://localhost:8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
