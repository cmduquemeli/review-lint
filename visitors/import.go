package visitors

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/cmduquemeli/review-lint/inspector"
)

type (
	ImportHandler struct {
		packageImport string
		callback      importCallbackFunc
	}
	importCallbackFunc func(fset *token.FileSet, n *ast.ImportSpec, packageName string)
)

func NewImportHandler() *ImportHandler {
	return &ImportHandler{}
}

func (h *ImportHandler) WithPackageImport(packageImport string) *ImportHandler {
	h.packageImport = packageImport
	return h
}

func (h *ImportHandler) WithCalback(callback importCallbackFunc) *ImportHandler {
	h.callback = callback
	return h
}

func (h *ImportHandler) Register() *ImportHandler {
	inspector.RegisterHandler("*ast.ImportSpec", h)
	return h
}

func (h *ImportHandler) Process(fset *token.FileSet, n ast.Node) {
	node := n.(*ast.ImportSpec)
	if h.packageImport != "" && node.Path.Value == h.packageImport {
		packageName := ""
		if node.Name != nil {
			packageName = node.Name.Name
		} else {
			parts := strings.Split(node.Path.Value, "/")
			packageName = parts[len(parts)-1]
			packageName = packageName[:len(packageName)-1]
		}
		h.callback(fset, node, packageName)
	} else {
		h.callback(fset, node, "")
	}
}
