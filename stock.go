// TODO: fugo側でうまい具合にcast
package googleFinance

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

// Stock stores stock data
type Stock struct {
	Code          string    `json:"t"` // Depending on the market (f.g. integer code for TYO, ticker for NASDAQ...etc)
	Name          string    `json:"name"`
	Price         float64   `json:"l_fix,string"`
	Change        float64   `json:"c_fix,string"`
	ChangePercent float64   `json:"cp_fix,string"`
	UpdatedAt     time.Time `json:"lt_dts,string"`
}

// parseToStocks parses stock json data to its struct
func parseToStocks(stockJSON []byte) (*[]Stock, error) {
	s := bytes.NewReader(stockJSON)
	var newStockData *[]Stock
	dec := json.NewDecoder(s)
	dec.Decode(&newStockData)
	if newStockData == nil {
		return nil, errors.New("failed to parse stock")
	}
	return newStockData, nil
}
