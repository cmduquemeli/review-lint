package visitors

import (
	"go/ast"
	"go/token"

	"github.com/cmduquemeli/review-lint/inspector"
)

type (
	SelectorHandler struct {
		ident    string
		selName  string
		callback selectorCallbackFunc
	}
	selectorCallbackFunc func(fset *token.FileSet, n *ast.SelectorExpr)
)

func NewSelectorHandler() *SelectorHandler {
	return &SelectorHandler{}
}

func (h *SelectorHandler) WithIdent(ident string) *SelectorHandler {
	h.ident = ident
	return h
}

func (h *SelectorHandler) WithSelName(selName string) *SelectorHandler {
	h.selName = selName
	return h
}

func (h *SelectorHandler) WithCalback(callback selectorCallbackFunc) *SelectorHandler {
	h.callback = callback
	return h
}

func (h *SelectorHandler) Register() *SelectorHandler {
	inspector.RegisterHandler("*ast.SelectorExpr", h)
	return h
}

func (h *SelectorHandler) Process(fset *token.FileSet, n ast.Node) {
	node := n.(*ast.SelectorExpr)
	if ident, isIdent := node.X.(*ast.Ident); isIdent {
		if h.ident == ident.Name {
			if h.selName == node.Sel.Name {
				h.callback(fset, node)
			}
		}
	}
}
