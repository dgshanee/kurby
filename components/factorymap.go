package components

var factoryMap = map[string]func(cmp ComponentProp) Component{
	"h1":   h1Factory,
	"body": bodyFactory,
	"box":  boxFactory,
	"p":    pFactory,
}
