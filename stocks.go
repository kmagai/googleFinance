package googleFinance

import (
	"bytes"
	"fmt"

	"github.com/olekukonko/tablewriter"
)

// Stocks is an array of Stock
type Stocks []Stock

// String defines what we get on printing Stocks.
func (stks Stocks) String() string {
	var buf bytes.Buffer

	buf.WriteString("\n")
	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"CODE", "NAME", "PRICE", "CHANGE", "LAST_UPDATE"})
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_RIGHT)

	const layout = "2006-01-02 15:04:05"
	for _, gstk := range stks {
		table.Append([]string{
			gstk.Code,
			gstk.Name,
			fmt.Sprint(gstk.Price),
			fmt.Sprint(roundAt(gstk.Change, 2)) + " (" + fmt.Sprint(roundAt(gstk.ChangePercent, 2)) + "%)",
			gstk.UpdatedAt.Format(layout),
		})
	}
	table.Render()
	buf.WriteString("\n")
	return buf.String()
}
