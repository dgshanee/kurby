package components

import (
	"encoding/xml"
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

func (b Body) Render(v interface{}) string {
	return ""
}
