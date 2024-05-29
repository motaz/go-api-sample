package main

import (
	"bytes"
	"encoding/json"
	"os"
)

func storeOrder(Data ordertype) error {

	bytes, _ := json.Marshal(Data)
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err == nil {
		defer f.Close()

		_, err = f.WriteString(string(bytes) + "\n")
	}
	return err
}

func retreiveOrders() (orders []ordertype, err error) {

	var data []byte
	data, err = os.ReadFile("data.txt")
	if err == nil {
		for _, orderBytes := range bytes.Split(data, []byte("\n")) {
			if !bytes.Equal(orderBytes, []byte{}) {
				var order ordertype
				err = json.Unmarshal(orderBytes, &order)
				orders = append(orders, order)
			}
		}
	}
	return
}

func deleteOrder(phone string) (err error) {

	var data []byte
	data, err = os.ReadFile("data.txt")
	if err == nil {
		New := ""
		for _, orderBytes := range bytes.Split(data, []byte("\n")) {
			orderString := string(orderBytes)
			if orderString != "" {
				var order ordertype
				err := json.Unmarshal(orderBytes, &order)
				if err == nil && order.Phone != phone {
					New += orderString + "\n"
				}
			}
		}
		os.WriteFile("data.txt", []byte(New), 0)
	}
	return
}
