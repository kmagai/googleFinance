package googleFinance

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kmagai/fugo/lib"
)

const googleURL = "http://www.google.com/finance/info?infotype=infoquoteall&q=%s"

type API struct {
}

// GetStocks gets stock struct from google API
func (api *API) GetStocks(codes interface{}) (*[]fugo.Stock, error) {
	var query string
	switch st := codes.(type) {
	case string:
		query = st
	case []string:
		query = buildQuery(st)
	}

	res, err := http.Get(buildFetchURL(query))
	if err != nil {
		return nil, errors.New("failed to fetch")
	}
	defer res.Body.Close()

	stockJSON, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("couldn't properly read response. It could be a problem with a remote host")
	}

	return fugo.ParseToStocks(trimSlashes(stockJSON))
}

// trimSlashes trims useless slashes in Google Finance API response
func trimSlashes(json []byte) []byte {
	return []byte(string(json)[3:])
}

// buildFetchURL builds Google Finance API url with the query specified
func buildFetchURL(query string) string {
	return fmt.Sprintf(googleURL, query)
}

// buildQuery builds query specifically for Google Finance API with the stock
// TODO: changed interface!!!
func buildQuery(codes []string) string {
	var query string
	for _, code := range codes {
		query += code + ","
	}
	return query
}
