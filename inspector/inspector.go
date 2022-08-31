package inspector

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
)

type Handler interface {
	Process(fset *token.FileSet, node ast.Node)
}

var handlers = map[string][]Handler{}

func File(filePath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse file %s, err %v", filePath, err))
	}
	Inspect(fset, f)
}

func Inspect(fset *token.FileSet, node ast.Node) {
	ast.Inspect(node, func(n ast.Node) bool {
		if reflectType := reflect.TypeOf(n); reflectType != nil {
			if handlers, found := handlers[reflectType.String()]; found {
				for _, handler := range handlers {
					handler.Process(fset, n)
				}
			}
		}

		return true
	})
}

func RegisterHandler(nodeType string, handler Handler) {
	typeHandlers := []Handler{}
	if registeredHandlers, found := handlers[nodeType]; found {
		typeHandlers = registeredHandlers
	}
	typeHandlers = append(typeHandlers, handler)
	handlers[nodeType] = typeHandlers
}
