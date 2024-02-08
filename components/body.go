package components

import (
	"encoding/xml"
	"github.com/dgshanee/kurby/styles"
)

type Body struct {
	XMLName xml.Name `xml:"Body"`
	ComponentProp
}

func bodyFactory(cmp ComponentProp) Component {
	var t Body
	t.ComponentProp = cmp
	return t
}

<<<<<<< HEAD
func (b Body) Render(style styles.Styles, renderedChildren ...[]string) string {
=======
func (b Body) Render(v string, inline bool) string {
>>>>>>> boxes
	return ""
}
