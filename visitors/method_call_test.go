package visitors

import (
	"fmt"
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewCallHandler(t *testing.T) {
	assertions := require.New(t)

	handler := NewCallHandler()

	assertions.NotNil(handler)
}

func Test_CallHandler_WithIdent_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewCallHandler()
	expectedIdent := "abc"

	h := handler.WithIdent(expectedIdent)

	assertions.Equal(handler, h)
	assertions.Equal(expectedIdent, h.ident)
}

func Test_CallHandler_WithFunctionName_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewCallHandler()
	expectedFuncitionName := "abc"

	h := handler.WithFunctionName(expectedFuncitionName)

	assertions.Equal(handler, h)
	assertions.Equal(expectedFuncitionName, h.functionName)
}

func Test_CallHandler_WithCalback_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewCallHandler()
	var expectedCallback methodCallCallbackFunc
	expectedCallback = func(fset *token.FileSet, n *ast.CallExpr) {}

	h := handler.WithCalback(expectedCallback)

	assertions.Equal(handler, h)
	assertions.Equal(fmt.Sprintf("%v", expectedCallback), fmt.Sprintf("%v", h.callback))
}

func Test_CallHandler_Process_Callback(t *testing.T) {
	assertions := require.New(t)
	fileSet := &token.FileSet{}
	ident := "fmt"
	functionName := "Sprintf"
	node := &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: ident}, Sel: &ast.Ident{Name: functionName}}}

	called := false
	var callback methodCallCallbackFunc
	callback = func(fset *token.FileSet, n *ast.CallExpr) {
		if fset == fileSet && n == node {
			called = true
		}
	}
	handler := NewCallHandler().WithIdent(ident).WithFunctionName(functionName).WithCalback(callback)

	handler.Process(fileSet, node)

	assertions.True(called)
}

func Test_CallHandler_Process_NoCall_Callback(t *testing.T) {
	assertions := require.New(t)
	fileSet := &token.FileSet{}
	ident := "fmt"
	functionName := "Sprintf"
	node := &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: ident}, Sel: &ast.Ident{Name: functionName}}}

	called := false
	var callback methodCallCallbackFunc
	callback = func(fset *token.FileSet, n *ast.CallExpr) {
		if fset == fileSet && n == node {
			called = true
		}
	}
	handler := NewCallHandler().WithIdent(ident).WithFunctionName("functionName").WithCalback(callback)

	handler.Process(fileSet, node)

	assertions.False(called)
}
