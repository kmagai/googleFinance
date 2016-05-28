package main

import (
	"fmt"

	"github.com/kmagai/google_finance_api"
)

func main() {
	api := stock.NewGoogleAPI()
	stock, err := api.GetStock("6178") // Japan Post Holdings Co Ltd
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stock)
	return
}
