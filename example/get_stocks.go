package main

import (
	"fmt"

	"github.com/kmagai/googleFinance"
)

func main() {
	api := googleFinance.API{}
	codes := []string{"GOOG", "AAPL", "9984"} // 9984 for SoftBank Group Corp
	stocks, err := api.GetStocks(codes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stocks)
	return
}
