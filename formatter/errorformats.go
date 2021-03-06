package formatter

import (
	"fmt"

	"github.com/TomiLabo/onion/types"
)

func QuickFixPrintResult(filename string, r []types.Status) {
	for _, v := range r {
		fmt.Printf("%s|%d col %d| %s\n", filename, v.Line, v.Column, v.Message)
	}
}
