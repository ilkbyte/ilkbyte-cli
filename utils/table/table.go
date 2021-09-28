package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func Create(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
