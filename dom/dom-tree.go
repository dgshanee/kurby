package dom

import (
	"cs/kurby/components"
	"encoding/xml"
	"fmt"
	"strings"
)

type Node struct {
	component components.Component
	children  *[]Node
}

var (
	root Node
)

func BuildDOMTree(xmlFile []byte) {
	var cmp components.Body
	err := xml.Unmarshal(xmlFile, &cmp)
	if err != nil {
		panic(err)
	}

	root.component = cmp
	root.children = &[]Node{}
	seeds := cmp.GetChildren()
	buildDOMTree(seeds, &root)
	fmt.Println("Done building")
	displayTree(root, 0)
}

func buildDOMTree(cmps []components.ComponentProp, parent *Node) {
	for _, c := range cmps {
		comp, ok := components.BuildComponent(c)
		if !ok {
			panic("")
		}
		var newNode = Node{
			component: comp,
			children:  &[]Node{},
		}
		if parent.children == nil {
			parent.children = &[]Node{}
		}
		*parent.children = append(*parent.children, newNode)
		buildDOMTree(comp.GetChildren(), &newNode)
	}

}

func displayTree(root Node, depth int) {
	indent := strings.Repeat("   ", depth)

	fmt.Printf("%sComponent:%s\n", indent, root.component.GetInnerText())

	for _, child := range *root.children {
		displayTree(child, depth+1)
	}
}
