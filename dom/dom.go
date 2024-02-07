package dom

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/dgshanee/kurby/components"
	"github.com/pterm/pterm"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Display() {
	xmlFile, err := os.ReadFile("structure.kurby")
	check(err)

	rootNode := BuildDOMTree(xmlFile)

	displayNode(rootNode)
}

func displayNode(node Node) {
	if node.component.GetXMLName() == "box" {
		var rendered []string
		for _, i := range *node.children {
			rendered = append(rendered, i.component.Render("", false))
		}

		pterm.Print(node.component.Render("", false))
	} else {

		pterm.Print(node.component.Render("", false))
		for _, i := range *node.children {
			displayNode(i)
		}
	}
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
