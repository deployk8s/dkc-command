package prepare

import (
	"strings"

	"github.com/deployKubernetesInCHINA/dkc-command/src"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"

	"github.com/alexeyco/simpletable"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/inventory"
)

func embarkData(t *src.Table, tds []inventory.Host) {

	for _, row := range tds {
		h := inventory.NewHostInstance(row)
		if err := h.Connect(); err == nil {
			defer h.Close()
			row.Arch, err = h.CombinedOutput("cat /etc/redhat-release")
			if err != nil {
				log.Log.Println(err.Error())
			} else {
				row.Arch = strings.ReplaceAll(row.Arch, "\n", "")
			}
		} else {
			log.Log.Error(err.Error())
		}
		var rma, rn, rsg, rom string

		if row.IsMaster {
			rma = "yes"
		}
		if row.IsNode {
			rn = "yes"
		}
		if row.IsOm {
			rom = "yes"
		}
		//if row.IsTracking {
		//	rtra = "yes"
		//}
		if row.IsSg {
			rsg = "yes"
		}

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: row.Hostname},
			{Align: simpletable.AlignCenter, Text: rma},
			{Align: simpletable.AlignCenter, Text: rn},
			{Align: simpletable.AlignCenter, Text: row.IsMongo},
			{Align: simpletable.AlignCenter, Text: rom},
			//{Text: rtra},
			{Align: simpletable.AlignCenter, Text: rsg},
			{Align: simpletable.AlignCenter, Text: row.Ip},
			{Align: simpletable.AlignCenter, Text: row.Arch},
		}

		t.Body.Cells = append(t.Body.Cells, r)
	}
}
