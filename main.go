package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ordertype struct {
	Phone   string
	Name    string
	Address string
}

type Response struct {
	Message string
	Success bool
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	inputBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		var input ordertype
		err := json.Unmarshal(inputBytes, &input)
		var output Response
		if err == nil {
			err = Add(input)
		}
		if err != nil {
			output.Message = err.Error()
		} else {
			output.Success = true
		}

		data, _ := json.Marshal(output)
		w.Write(data)
	}

}

func RemoveOrder(w http.ResponseWriter, r *http.Request) {
	inputBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		input := string(inputBytes)
		var output Response
		if err == nil {
			err = Remove(input)
		}
		if err != nil {
			output.Message = err.Error()
		} else {
			output.Success = true
		}

		data, _ := json.Marshal(output)
		w.Write(data)
	}

}

type GetResponse struct {
	Message string
	Success bool
	Orders  []ordertype
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var output GetResponse
	orders, err := Get()
	if err == nil {
		output.Orders = orders
		output.Success = true
	} else {
		output.Message = err.Error()
	}
	data, _ := json.Marshal(output)
	w.Write(data)
}

func main() {
	http.HandleFunc("/addorder", AddOrder)
	http.HandleFunc("/removeorder", RemoveOrder)
	http.HandleFunc("/getorders", GetOrders)
	http.ListenAndServe(":8888", nil)
}
