package components

import (
	"github.com/pterm/pterm"
)

type Box struct {
	ComponentProp
}

func boxFactory(cmp ComponentProp) Component {
	var t Box
	t.ComponentProp = cmp
	return t
}

func (c Box) Render(v string, inline bool) string {
	box := pterm.DefaultBox

	var panels [][]pterm.Panel
	var curr []pterm.Panel

	if c.ComponentProp.Title != "" {
		box = *box.WithTitle(c.ComponentProp.Title)
	}
	if renderPanels != "" {
		return box.Sprint(renderPanels)
	}
	return box.Sprint(c.InnerText)
}
