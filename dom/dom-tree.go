package dom

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"

	"github.com/dgshanee/kurby/components"
)

type Node struct {
	component components.Component
	children  *[]Node
}

var (
	root Node
)

func BuildDOMTree(xmlFile []byte) Node {
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
	return root
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

	fmt.Printf("%sComponent:%s, type %s\n", indent, root.component.GetInnerText(), reflect.TypeOf(root.component))

	for _, child := range *root.children {
		displayTree(child, depth+1)
	}
}
