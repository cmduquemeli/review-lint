package validators

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/cmduquemeli/review-lint/visitors"
)

type (
	MockAnythingValidator struct {
		packageImport string
		packageName   string
		failFunc      validatorFailFunc
	}
)

func NewMockAnythingValidator() *MockAnythingValidator {
	return &MockAnythingValidator{packageImport: `"github.com/stretchr/testify/mock"`, failFunc: func() {}}
}

func (v *MockAnythingValidator) WithPackageImport(packageImport string) *MockAnythingValidator {
	v.packageImport = packageImport
	return v
}

func (v *MockAnythingValidator) WithFailFunc(failFunc validatorFailFunc) *MockAnythingValidator {
	v.failFunc = failFunc
	return v
}

func (v *MockAnythingValidator) Register() *MockAnythingValidator {
	sh := visitors.NewSelectorHandler().
		WithCalback(func(fset *token.FileSet, n *ast.SelectorExpr) {
			fmt.Println("Use concrete value instead of mock.Anything", fset.Position(n.Pos()).String())
			v.failFunc()
		}).
		WithSelName("Anything").
		Register()

	visitors.NewImportHandler().
		WithPackageImport(v.packageImport).
		WithCalback(func(_ *token.FileSet, _ *ast.ImportSpec, packageName string) {
			if packageName != "" && v.packageName == "" {
				v.packageName = packageName
				sh.WithIdent(packageName)
			}
		}).
		Register()

	return v
}
