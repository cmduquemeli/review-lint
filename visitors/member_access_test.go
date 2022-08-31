package visitors

import (
	"fmt"
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewSelectorHandler(t *testing.T) {
	assertions := require.New(t)

	handler := NewSelectorHandler()

	assertions.NotNil(handler)
}

func Test_SelectorHandler_WithIdent_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewSelectorHandler()
	expectedIdent := "abc"

	h := handler.WithIdent(expectedIdent)

	assertions.Equal(handler, h)
	assertions.Equal(expectedIdent, h.ident)
}

func Test_SelectorHandler_WithSelName_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewSelectorHandler()
	expectedSelName := "abc"

	h := handler.WithSelName(expectedSelName)

	assertions.Equal(handler, h)
	assertions.Equal(expectedSelName, h.selName)
}

func Test_SelectorHandler_WithCalback_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewSelectorHandler()
	var expectedCallback selectorCallbackFunc
	expectedCallback = func(fset *token.FileSet, n *ast.SelectorExpr) {}

	h := handler.WithCalback(expectedCallback)

	assertions.Equal(handler, h)
	assertions.Equal(fmt.Sprintf("%v", expectedCallback), fmt.Sprintf("%v", h.callback))
}

func Test_SelectorHandler_Process_Callback(t *testing.T) {
	assertions := require.New(t)
	fileSet := &token.FileSet{}
	ident := "mock"
	selName := "Anything"
	node := &ast.SelectorExpr{X: &ast.Ident{Name: ident}, Sel: &ast.Ident{Name: selName}}
	called := false
	var callback selectorCallbackFunc
	callback = func(fset *token.FileSet, n *ast.SelectorExpr) {
		if fset == fileSet && n == node {
			called = true
		}
	}
	handler := NewSelectorHandler().WithIdent(ident).WithSelName(selName).WithCalback(callback)

	handler.Process(fileSet, node)

	assertions.True(called)
}

func Test_SelectorHandler_Process_NoCall_Callback(t *testing.T) {
	assertions := require.New(t)
	fileSet := &token.FileSet{}
	ident := "mock"
	selName := "Anything"
	node := &ast.SelectorExpr{X: &ast.Ident{Name: ident}, Sel: &ast.Ident{Name: selName}}
	called := false
	var callback selectorCallbackFunc
	callback = func(fset *token.FileSet, n *ast.SelectorExpr) {
		if fset == fileSet && n == node {
			called = true
		}
	}
	handler := NewSelectorHandler().WithIdent(ident).WithSelName("selName").WithCalback(callback)

	handler.Process(fileSet, node)

	assertions.False(called)
}
