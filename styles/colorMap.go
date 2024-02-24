package styles

import "github.com/pterm/pterm"

var (
	ColorMap map[string]*pterm.Style = map[string]*pterm.Style{
		"cyan": pterm.FgCyan.ToStyle(),
	}
)
