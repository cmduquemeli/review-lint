package validators

import (
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/imports"
)

type validatorFailFunc func()

func parseSource(src string) (*token.FileSet, ast.Node) {
	imp, _ := imports.Process("", []byte(src), nil)
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", imp, 0)
	return fset, f
}
