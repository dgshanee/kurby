package components

import (
	"encoding/xml"
)

type Component interface {
	Render(v string, inline bool) string
	GetChildren() []ComponentProp
	GetInnerText() string
	GetXMLName() string
}

type ComponentProp struct {
	XMLName   xml.Name
	Height    int             `xml:"height,attr"`
	Width     int             `xml:"width,attr"`
	InnerText string          `xml:",chardata"`
	Children  []ComponentProp `xml:",any"`
	Inline    bool            `xml:"inline,attr"`
	Title     string          `xml:"title,attr"`
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
