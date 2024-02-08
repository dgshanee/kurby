package components

import (
	"encoding/xml"
	"strings"

	"github.com/dgshanee/kurby/styles"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
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

<<<<<<< HEAD
func (c h1) Render(style styles.Styles, renderedChildren ...[]string) string {
=======
func (c h1) Render(v string, inline bool) string {
>>>>>>> boxes
	var res string
	letters := putils.LettersFromString(strings.TrimSpace(c.GetInnerText()))

	res, err := pterm.DefaultBigText.WithLetters(
		letters,
	).Srender()

	if err != nil {
		return "ERROR"
	}

	return res
}
