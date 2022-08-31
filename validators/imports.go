package validators

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"strconv"

	"github.com/cmduquemeli/review-lint/visitors"
)

type (
	ImportsValidator struct {
		whitelistImports []*regexp.Regexp
		blacklistImports []*regexp.Regexp
		failFunc         validatorFailFunc
	}
)

func NewImportsValidator() *ImportsValidator {
	return &ImportsValidator{blacklistImports: []*regexp.Regexp{}, failFunc: func() {}}
}

func (v *ImportsValidator) WithBlacklist(blacklist []string) *ImportsValidator {
	for _, blockedImport := range blacklist {
		v.blacklistImports = append(v.blacklistImports, regexp.MustCompile("^"+blockedImport+"$"))
	}
	return v
}

func (v *ImportsValidator) WithWhitelist(whitelist []string) *ImportsValidator {
	for _, knownImport := range whitelist {
		v.whitelistImports = append(v.whitelistImports, regexp.MustCompile("^"+knownImport+"$"))
	}
	return v
}

func (v *ImportsValidator) WithFailFunc(failFunc validatorFailFunc) *ImportsValidator {
	v.failFunc = failFunc
	return v
}

func (v *ImportsValidator) Register() *ImportsValidator {
	visitors.NewImportHandler().
		WithCalback(func(fset *token.FileSet, node *ast.ImportSpec, _ string) {
			importName, _ := strconv.Unquote(node.Path.Value)
			if !v.validateBlockedImport(fset, node, importName) {
				v.validateUnknownImport(fset, node, importName)
			}
		}).
		Register()

	return v
}

func (v *ImportsValidator) validateBlockedImport(fset *token.FileSet, node *ast.ImportSpec, importName string) bool {
	for _, regex := range v.blacklistImports {
		if regex.Match([]byte(importName)) {
			fmt.Printf("Import \"%s\" is not allowed - %s\r\n", importName, fset.Position(node.Pos()))
			v.failFunc()
			return true
		}
	}
	return false
}

func (v *ImportsValidator) validateUnknownImport(fset *token.FileSet, node *ast.ImportSpec, importName string) bool {
	for _, regex := range v.whitelistImports {
		if regex.Match([]byte(importName)) {
			return false
		}
	}
	fmt.Printf("Import \"%s\" is not known - %s\r\n", importName, fset.Position(node.Pos()))
	v.failFunc() // TODO Fail when is unknown?
	return true
}
