package components

import (
	"encoding/xml"
)

type Body struct {
	XMLName xml.Name `xml:"Body"`
	ComponentProp
}
