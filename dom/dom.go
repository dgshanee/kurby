package dom

import (
	"cs/kurby/components"
	"encoding/xml"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Display() {
	xmlFile, err := os.ReadFile("structure.kurby")
	check(err)

	BuildDOMTree(xmlFile)
}

func GetElements(xmlSlice []byte) {
	var root components.ComponentProp
	xml.Unmarshal(xmlSlice, &root)
	for _, chi := range root.GetChildren() {
		fmt.Println("Parent:", chi)
		children := chi.GetChildren()
		if len(children) > 0 {
			fmt.Println("Children:")
			for _, c := range children {
				fmt.Println(c)
			}
		}
	}
}
