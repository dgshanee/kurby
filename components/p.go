package components

import (
	"github.com/pterm/pterm"
)

type P struct {
	ComponentProp
}

func pFactory(cmp ComponentProp) Component {
	var t P
	t.ComponentProp = cmp
	return t
}

func (c P) Render(v string, inline bool) string {
	return pterm.DefaultBasicText.Sprint(c.InnerText)
}
