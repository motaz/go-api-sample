package main

type Response struct {
	Message string
	Success bool
}

func Add(input ordertype) (output Response) {
	if input.Address == "" || input.Address == "" || input.Address == "" {
		output.Message = "empty param"
	} else {
		err := orderAdd(input)
		if err != nil {
			output.Message = err.Error()
		} else {
			output.Success = true
		}
	}
	return
}

func Remove(phone string) (output Response) {
	if err := orderRemove(phone); err != nil {
		output.Message = err.Error()
	} else {
		output.Success = true
	}
	return
}

type GetResponse struct {
	Message string
	Success bool
	Orders  []ordertype
}

func Get() (output GetResponse) {
	orders, err := getOrders()
	if err != nil {
		output.Message = err.Error()
	} else {
		output.Success = true
		output.Orders = orders
	}
	return
}
