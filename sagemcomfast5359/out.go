package sagemcomfast5359

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

const padding = 3

// table prints a table.
func Table(output io.Writer, devices []Device) {
	tw := tabwriter.NewWriter(output, 0, 0, padding, ' ', 0)
	columns := []string{"IP", "PhysAddress", "Name"}
	fmt.Fprintln(tw, strings.Join(columns, "\t"))
	for _, d := range devices {
		fmt.Fprintf(tw, "%s\t%s\t%s\t\n", d.IPAddress, d.PhysAddress, d.Name)
	}
	tw.Flush()
}
