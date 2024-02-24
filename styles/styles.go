package styles

import "fmt"

type Styles struct {
	Inline        bool   `xml:"inline,attr"`
	Height        int    `xml:"height,attr"`
	Width         int    `xml:"width,attr"`
	Padding       int    `xml:"padding,attr"`
	BottomPadding int    `xml:"bottom-padding,attr"`
	RGB           string `xml:"rgb,attr"`
}

func GetRGBFromString(s string) (uint8, uint8, uint8) {
	var r, g, b uint8

	_, err := fmt.Sscanf(s, "%d,%d,%d", &r, &g, &b)
	if err != nil {
		fmt.Println("Error parsing RGB")
		fmt.Println(err)
		return 0, 0, 0
	}

	return r, g, b
}
