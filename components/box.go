package components

import (
	"fmt"
	"github.com/dgshanee/kurby/styles"
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

func (c Box) Render(style styles.Styles, renderedChildren ...[]string) string {
	box := pterm.DefaultBox

	var WIDTH int = 500

	if style.Width != 0 {
		WIDTH = style.Width
	}

	var panels [][]pterm.Panel
	var curr []pterm.Panel
	var renderPanels string

	for _, renderedChild := range renderedChildren {
		for _, str := range renderedChild {
			currPanel := pterm.Panel{Data: str}
			curr = append(curr, currPanel)

			if len(curr) >= WIDTH {
				panels = append(panels, curr)
				curr = []pterm.Panel{}
			}
		}
		if len(curr) > 0 {
			panels = append(panels, curr)
		}
	}

	finPanels := pterm.Panels(panels)
	renderPanels, _ = pterm.DefaultPanel.WithPanels(finPanels).Srender()

	if style.Padding > 0 {
		box = *box.WithTopPadding(style.Padding)
		box = *box.WithLeftPadding(style.Padding)
		box = *box.WithRightPadding(style.Padding)
		box = *box.WithBottomPadding(style.Padding)
	}

	if style.BottomPadding > 0 {
		box = *box.WithBottomPadding(style.BottomPadding)
	}

	if c.ComponentProp.Title != "" {
		box = *box.WithTitle(c.ComponentProp.Title)
	}
	if renderPanels != "" {
		return box.Sprint(renderPanels)
	}
	return box.Sprint(c.InnerText)
}
