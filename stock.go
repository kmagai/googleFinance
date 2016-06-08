package googleFinance

import (
	"bytes"
	"fmt"
	"math"
	"time"

	"github.com/olekukonko/tablewriter"
)

// Stock stores stock data from Google Finance API.
type Stock struct {
	Code          string    `json:"t"` // Depending on the market (f.g. integer code for TYO, ticker for NASDAQ...etc)
	Name          string    `json:"name"`
	Price         float64   `json:"l_fix,string"`
	Change        float64   `json:"c_fix,string"`
	ChangePercent float64   `json:"cp_fix,string"`
	UpdatedAt     time.Time `json:"lt_dts,string"`
}

// GetCode is getter method for Stock's Code.
func (stk *Stock) GetCode() string {
	return stk.Code
}

// String defines what we get on printing Stock.
func (stk Stock) String() string {
	var buf bytes.Buffer

	buf.WriteString("\n")
	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"CODE", "NAME", "PRICE", "CHANGE", "LAST_UPDATE"})
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_RIGHT)

	const layout = "2006-01-02 15:04:05"
	table.Append([]string{
		stk.Code,
		stk.Name,
		fmt.Sprint(stk.Price),
		fmt.Sprint(roundAt(stk.Change, 2)) + " (" + fmt.Sprint(roundAt(stk.ChangePercent, 2)) + "%)",
		stk.UpdatedAt.Format(layout),
	})
	table.Render()
	buf.WriteString("\n")
	return buf.String()
}

// ToStocks takes Stock slices and parses it into a Stocks struct
func (stk *Stock) ToStocks(stks []Stock) Stocks {
	return Stocks(stks)
}

func roundAt(f float64, roundAt int) float64 {
	shift := math.Pow(10, float64(roundAt))
	return round(f*shift) / shift
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}
