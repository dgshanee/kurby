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

<<<<<<< HEAD
func (c P) Render(style styles.Styles, renderedChildren ...[]string) string {
=======
func (c P) Render(v string, inline bool) string {
>>>>>>> boxes
	return pterm.DefaultBasicText.Sprint(c.InnerText)
}
