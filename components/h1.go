package components

import (
	"encoding/xml"
	"strings"

	"github.com/pterm/pterm"
)

type h1 struct {
	XMLName xml.Name `xml:"h1"`
	ComponentProp
}

func h1Factory(cmp ComponentProp) Component {
	var t h1
	t.ComponentProp = cmp
	return t
}

func (c h1) Render() string {
	var res string
	if c.Inline {
		res = pterm.DefaultBasicText.Sprint(strings.TrimSpace(c.GetInnerText()))
	} else {
		res = pterm.DefaultBasicText.Sprintf(strings.TrimSpace(c.GetInnerText()))
	}
	return res
}
