package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/addorder", AddOrder)
	http.HandleFunc("/removeorder", RemoveOrder)
	http.HandleFunc("/getorders", GetOrders)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
