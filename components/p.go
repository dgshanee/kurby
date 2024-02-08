package components

import (
	"github.com/dgshanee/kurby/styles"
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

func (c P) Render(style styles.Styles, renderedChildren ...[]string) string {
	return pterm.DefaultBasicText.Sprint(c.InnerText)
}
