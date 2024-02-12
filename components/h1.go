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

func (c h1) Render(style styles.Styles, renderedChildren ...[]string) string {
	var res string
	str := strings.TrimSpace(c.GetInnerText())
	letters := putils.LettersFromString(str)
	if style.RGB != "" {
		r, g, b := styles.GetRGBFromString(style.RGB)

		letters = putils.LettersFromStringWithRGB(str, pterm.NewRGB(r, g, b))
	}

	res, err := pterm.DefaultBigText.WithLetters(
		letters,
	).Srender()

	if err != nil {
		return "ERROR"
	}

	return res
}
