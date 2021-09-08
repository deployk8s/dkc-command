package src

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

type Table struct {
	*simpletable.Table
}

func NewTable(cells ...string) *Table {
	var cs []*simpletable.Cell
	for _, v := range cells {
		cs = append(cs, &simpletable.Cell{Align: simpletable.AlignCenter, Text: v})
	}
	t := simpletable.New()
	t.Header = &simpletable.Header{
		Cells: cs,
		//Cells: []*simpletable.Cell{
		//	{Align: simpletable.AlignCenter, Text: "HOSTNAME"},
		//	{Align: simpletable.AlignCenter, Text: "MASTER"},
		//	{Align: simpletable.AlignCenter, Text: "NODE"},
		//	{Align: simpletable.AlignCenter, Text: "MONGO"},
		//	{Align: simpletable.AlignCenter, Text: "OM"},
		//	{Align: simpletable.AlignCenter, Text: "SG"},
		//	{Align: simpletable.AlignCenter, Text: "IP"},
		//	{Align: simpletable.AlignCenter, Text: "Release"},
		//},
	}
	return &Table{t}
}

func (t *Table) Show() {
	t.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(t.String())
}
