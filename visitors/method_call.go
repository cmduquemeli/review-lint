package visitors

import (
	"go/ast"
	"go/token"

	"github.com/cmduquemeli/review-lint/inspector"
)

type (
	CallHandler struct {
		ident        string
		functionName string
		callback     methodCallCallbackFunc
	}
	methodCallCallbackFunc func(fset *token.FileSet, n *ast.CallExpr)
)

func NewCallHandler() *CallHandler {
	return &CallHandler{}
}

func (h *CallHandler) WithIdent(ident string) *CallHandler {
	h.ident = ident
	return h
}

func (h *CallHandler) WithFunctionName(functionName string) *CallHandler {
	h.functionName = functionName
	return h
}

func (h *CallHandler) WithCalback(callback methodCallCallbackFunc) *CallHandler {
	h.callback = callback
	return h
}

func (h *CallHandler) Register() *CallHandler {
	inspector.RegisterHandler("*ast.CallExpr", h)
	return h
}

func (h *CallHandler) Process(fset *token.FileSet, n ast.Node) {
	node := n.(*ast.CallExpr)
	selector, isSelector := node.Fun.(*ast.SelectorExpr)
	if isSelector {
		ident, isIdent := selector.X.(*ast.Ident)
		if isIdent {
			packageName := ident.Name
			if packageName == h.ident {
				if h.functionName != "" {
					if h.functionName == selector.Sel.Name {
						h.callback(fset, node)
					}
				} else {
					h.callback(fset, node)
				}
			}
		}
	}
}
