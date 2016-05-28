package main

import (
	"fmt"

	"github.com/kmagai/googleFinance"
)

func main() {
	api := googleFinance.API{}
	stock, err := api.GetStocks("6178") // Japan Post Holdings Co Ltd
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stock)
	return
}
