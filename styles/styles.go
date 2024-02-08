package styles

type Styles struct {
	Inline        bool `xml:"inline,attr"`
	Height        int  `xml:"height,attr"`
	Width         int  `xml:"width,attr"`
	Padding       int  `xml:"padding,attr"`
	BottomPadding int  `xml:"bottom-padding,attr"`
}
