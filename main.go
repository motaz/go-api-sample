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

func AddOrder(w http.ResponseWriter, r *http.Request) {
	inputBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		var input ordertype
		err := json.Unmarshal(inputBytes, &input)
		var output Response
		if err == nil {
			output = Add(input)
		} else {
			output.Message = err.Error()
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
			output = Remove(input)
		} else {
			output.Message = err.Error()
		}

		data, _ := json.Marshal(output)
		w.Write(data)
	}

}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	output := Get()
	data, _ := json.Marshal(output)
	w.Write(data)
}

func main() {
	http.HandleFunc("/addorder", AddOrder)
	http.HandleFunc("/removeorder", RemoveOrder)
	http.HandleFunc("/getorders", GetOrders)
	http.ListenAndServe(":8888", nil)
}
