package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Message string
	Success bool
}

type GetResponse struct {
	Message string
	Success bool
	Orders  []ordertype
}

type ordertype struct {
	Phone   string
	Name    string
	Address string
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		var request ordertype
		err := json.Unmarshal(requestBytes, &request)
		var response Response
		if err == nil {
			if request.Address == "" || request.Address == "" || request.Address == "" {
				response.Message = "empty param"
			} else {
				err := orderAdd(request)
				if err != nil {
					response.Message = err.Error()
				} else {
					response.Success = true
				}
			}
		} else {
			response.Message = err.Error()
		}

		data, _ := json.Marshal(response)
		w.Write(data)
	}

}

func RemoveOrder(w http.ResponseWriter, r *http.Request) {
	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		phone := string(requestBytes)
		var response Response
		if err == nil {
			if err := orderRemove(phone); err != nil {
				response.Message = err.Error()
			} else {
				response.Success = true
			}
		} else {
			response.Message = err.Error()
		}

		data, _ := json.Marshal(response)
		w.Write(data)
	}

}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := getOrders()
	var response GetResponse
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Success = true
		response.Orders = orders
	}
	data, _ := json.Marshal(response)
	w.Write(data)
}
