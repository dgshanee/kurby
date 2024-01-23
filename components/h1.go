package components

import "encoding/xml"

type h1 struct {
	XMLName xml.Name `xml:"h1"`
	ComponentProp
}

func h1Factory(cmp ComponentProp) Component {
	var t h1
	t.ComponentProp = cmp
	return t
}
