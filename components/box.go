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

func (c Box) Render(v interface{}) string {
	box := pterm.DefaultBox
	var renderPanels string
	var panels []pterm.Panel
	if strs, ok := v.([]string); ok {
		for _, str := range strs {
			panel := pterm.Panel{Data: str}
			panels = append(panels, panel)
		}

		finPanels := pterm.Panels{panels}

		renderPanels, _ = pterm.DefaultPanel.WithPanels(finPanels).Srender()
	}
	if c.ComponentProp.Title != "" {
		box = *box.WithTitle(c.ComponentProp.Title)
	}
	if renderPanels != "" {
		return box.Sprint(renderPanels)
	}
	return box.Sprint(c.InnerText)
}
