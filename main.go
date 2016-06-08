package googleFinance

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kmagai/fugo/lib/plugin"
)

const googleURL = "http://www.google.com/finance/info?infotype=infoquoteall&q=%s"

// API is the interface for the google finance resource
type API struct{}

// GetStock gets stock struct from google API.
func (api *API) GetStock(code string) (*Stock, error) {
	res, err := http.Get(buildFetchURL(code))
	if err != nil {
		return nil, errors.New("failed to fetch")
	}
	defer res.Body.Close()

	stockJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("couldn't properly read response. It could be a problem with a remote host")
	}

	return parseStock(trimSlashes(stockJSON))
}

// GetStocks gets stocks struct from google API.
func (api API) GetStocks(codes []string) (*Stocks, error) {
	query := buildQuery(codes)
	res, err := http.Get(buildFetchURL(query))
	if err != nil {
		return nil, errors.New("failed to fetch")
	}
	defer res.Body.Close()

	stockJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("couldn't properly read response. It could be a problem with a remote host")
	}

	return parseStocks(trimSlashes(stockJSON))
}

// GetStocker wraps the Stock into fugo Stocker interface
func (api API) GetStocker(code string) (plugin.Stocker, error) {
	stock, err := api.GetStock(code)
	if err != nil {
		return nil, err
	}
	return *toStocker(stock), nil
}

// GetStockers wraps the Stocks into fugo Stockers interface
func (api API) GetStockers(codes []string) (plugin.Stockers, error) {
	stocks, err := api.GetStocks(codes)
	if err != nil {
		return nil, err
	}
	stocker := toStockers(stocks)
	if stocker == nil {
		return nil, errors.New("failed to parse stock")
	}
	return *stocker, nil
}

// trimSlashes trims useless slashes in Google Finance API response.
func trimSlashes(json []byte) []byte {
	return []byte(string(json)[3:])
}

// buildFetchURL builds Google Finance API url with the query specified.
func buildFetchURL(query string) string {
	return fmt.Sprintf(googleURL, query)
}

// buildQuery builds query specifically for Google Finance API with the stock.
func buildQuery(codes []string) string {
	var query string
	for _, code := range codes {
		query += code + ","
	}
	return query
}

// ParseStock parses stock json into a Stock struct pointer.
func parseStock(JSON []byte) (*Stock, error) {
	s := bytes.NewReader(JSON)
	var gstks *Stocks
	dec := json.NewDecoder(s)
	err := dec.Decode(&gstks)
	if err != nil {
		return nil, err
	}
	if gstks == nil {
		return nil, errors.New("failed to parse stock")
	}
	if len(*gstks) > 1 {
		return nil, errors.New("got more than one stock data")
	}
	gstk := (*gstks)[0]
	return &gstk, nil
}

// ParseStocks parses stocks json into a Stocks struct pointer.
func parseStocks(JSON []byte) (*Stocks, error) {
	s := bytes.NewReader(JSON)
	var gstks *Stocks
	dec := json.NewDecoder(s)
	err := dec.Decode(&gstks)
	if err != nil {
		return nil, err
	}
	if gstks == nil {
		return nil, errors.New("failed to parse stock")
	}
	return gstks, nil
}

func toStocker(stck interface{}) *plugin.Stocker {
	stkr, ok := stck.(plugin.Stocker)
	if !ok {
		return nil
	}
	return &stkr
}

func toStockers(stcks interface{}) *plugin.Stockers {
	stkrs, ok := stcks.(plugin.Stockers)
	if !ok {
		return nil
	}
	return &stkrs
}
