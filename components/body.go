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

func (b Body) Render(style styles.Styles, renderedChildren ...[]string) string {
	return ""
}
