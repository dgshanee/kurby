package dom

import (
	"encoding/xml"
	"fmt"
	"github.com/dgshanee/kurby/components"
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

	rootNode := BuildDOMTree(xmlFile)

	displayNode(rootNode)
}

func displayNode(node Node) {
	fmt.Print(node.component.Render())
	for _, i := range *node.children {
		displayNode(i)
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
