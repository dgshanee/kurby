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
	component := node.component
	componentStyles := component.GetStyles()
	if component.GetXMLName() == "box" {
		var rendered []string
		for _, i := range *node.children {
			child := i.component
			childStyles := component.GetStyles()
			rendered = append(rendered, child.Render(childStyles))
		}
		pterm.Print(component.Render(componentStyles, rendered))
	} else {

		pterm.Print(component.Render(componentStyles))
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
