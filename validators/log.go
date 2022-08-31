package validators

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"strconv"
	"strings"

	"github.com/cmduquemeli/review-lint/visitors"
)

type (
	LogValidator struct {
		packageImport string
		packageName   string
		functionName  string
		failFunc      validatorFailFunc
	}
)

var (
	regexStr      = `^[a-z_]+$`
	regex         = regexp.MustCompile(regexStr)
	reservedWords = map[string]struct{}{
		"msg": {},
	}
)

func NewLogValidator() *LogValidator {
	return &LogValidator{packageImport: `"github.com/mercadolibre/fury_go-core/pkg/log"`, failFunc: func() {}}
}

func (v *LogValidator) WithPackageImport(packageImport string) *LogValidator {
	v.packageImport = packageImport
	return v
}

func (v *LogValidator) WithFunctionName(functionName string) *LogValidator {
	v.functionName = functionName
	return v
}

func (v *LogValidator) WithFailFunc(failFunc validatorFailFunc) *LogValidator {
	v.failFunc = failFunc
	return v
}

func (v *LogValidator) Register() *LogValidator {
	ch := visitors.NewCallHandler().
		WithCalback(func(fset *token.FileSet, n *ast.CallExpr) {
			v.StringArgumentNotCompliant(fset, n)
		}).
		Register()

	visitors.NewImportHandler().
		WithPackageImport(v.packageImport).
		WithCalback(func(_ *token.FileSet, _ *ast.ImportSpec, packageName string) {
			if packageName != "" && v.packageName == "" {
				v.packageName = packageName
				ch.WithIdent(packageName)
			}
		}).
		Register()

	return v
}

func (v *LogValidator) StringArgumentNotCompliant(fset *token.FileSet, n *ast.CallExpr) {
	for _, arg := range n.Args {
		if basicLit, isBasicLic := arg.(*ast.BasicLit); isBasicLic {
			if basicLit.Kind == token.STRING {
				strValue, _ := strconv.Unquote(basicLit.Value)
				if !regex.Match([]byte(strValue)) {
					fmt.Printf("Log string \"%s\" isn't compliant with regex \"%s\" - %s\r\n", strValue, regexStr, fset.Position(basicLit.Pos()))
					v.failFunc()
					return
				}
				if isReservedWord(strValue) {
					fmt.Printf("Log string \"%s\" is a reserved word - %s", strValue, fset.Position(basicLit.Pos()))
					v.failFunc()
					return
				}
			}
		}
	}
}

func isReservedWord(strValue string) bool {
	_, isReservedWord := reservedWords[strings.ToLower(strValue)]
	return isReservedWord
}
