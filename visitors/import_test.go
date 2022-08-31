package visitors

import (
	"fmt"
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewImportHandler(t *testing.T) {
	assertions := require.New(t)

	handler := NewImportHandler()

	assertions.NotNil(handler)
}

func Test_ImportHandler_WithPackageImport_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewImportHandler()
	expectedPackageImport := "abc"

	h := handler.WithPackageImport(expectedPackageImport)

	assertions.Equal(handler, h)
	assertions.Equal(expectedPackageImport, h.packageImport)
}

func Test_ImportHandler_WithCalback_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewImportHandler()
	var expectedCallback importCallbackFunc
	expectedCallback = func(fset *token.FileSet, n *ast.ImportSpec, packageName string) {}

	h := handler.WithCalback(expectedCallback)

	assertions.Equal(handler, h)
	assertions.Equal(fmt.Sprintf("%v", expectedCallback), fmt.Sprintf("%v", h.callback))
}

func Test_ImportHandler_ProcessWithPackage_Callback(t *testing.T) {
	assertions := require.New(t)
	fileSet := &token.FileSet{}
	packageImport := `"github.com/stretchr/testify/require"`
	node := &ast.ImportSpec{Path: &ast.BasicLit{Value: packageImport}}
	called := false
	var callback importCallbackFunc
	callback = func(fset *token.FileSet, n *ast.ImportSpec, packageName string) {
		if fset == fileSet && n == node && packageName == "require" {
			called = true
		}
	}
	handler := NewImportHandler().WithPackageImport(packageImport).WithCalback(callback)

	handler.Process(fileSet, node)

	assertions.True(called)
}

func Test_ImportHandler_ProcessWithPackageAlias_Callback(t *testing.T) {
	assertions := require.New(t)
	fileSet := &token.FileSet{}
	packageImport := `t "github.com/stretchr/testify/require"`
	node := &ast.ImportSpec{Name: &ast.Ident{Name: "t"}, Path: &ast.BasicLit{Value: packageImport}}
	called := false
	var callback importCallbackFunc
	callback = func(fset *token.FileSet, n *ast.ImportSpec, packageName string) {
		if fset == fileSet && n == node && packageName == "t" {
			called = true
		}
	}
	handler := NewImportHandler().WithPackageImport(packageImport).WithCalback(callback)

	handler.Process(fileSet, node)

	assertions.True(called)
}

func Test_ImportHandler_Process_WithoutPackageName_Callback(t *testing.T) {
	assertions := require.New(t)
	fileSet := &token.FileSet{}
	node := &ast.ImportSpec{}
	called := false
	var callback importCallbackFunc
	callback = func(fset *token.FileSet, n *ast.ImportSpec, packageName string) {
		if fset == fileSet && n == node && packageName == "" {
			called = true
		}
	}
	handler := NewImportHandler().WithCalback(callback)

	handler.Process(fileSet, node)

	assertions.True(called)
}
