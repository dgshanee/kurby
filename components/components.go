package components

import (
	"encoding/xml"

	"github.com/dgshanee/kurby/styles"
)

type Component interface {
	Render(style styles.Styles, renderedChildren ...[]string) string
	GetChildren() []ComponentProp
	GetInnerText() string
	GetXMLName() string
	GetStyles() styles.Styles
}

type ComponentProp struct {
	XMLName   xml.Name
	InnerText string          `xml:",chardata"`
	Children  []ComponentProp `xml:",any"`
	Title     string          `xml:"title,attr"`
	styles.Styles
}

func (cmp ComponentProp) GetStyles() styles.Styles {
	return cmp.Styles
}

func (cmp ComponentProp) isInline() bool {
	return cmp.Inline
}

func (cmp ComponentProp) GetInnerText() string {
	return cmp.InnerText
}

func (cmp ComponentProp) Render() string {
	return ""
}

func (cmp ComponentProp) GetChildren() []ComponentProp {

	return cmp.Children
}

func (cmp ComponentProp) GetXMLName() string {
	return string(cmp.XMLName.Local)
}

func BuildComponent(cmp ComponentProp) (Component, bool) {
	componentName := cmp.GetXMLName()

	factory, ok := factoryMap[componentName]
	if ok {
		return factory(cmp), true
	} else {
		return nil, false
	}
}
