package validators

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/cmduquemeli/review-lint/inspector"
	"github.com/stretchr/testify/require"
)

const srcTest = `
package main	

import "fmt"

var a = fmt.Sprintf("%s", "a")
`

func Test_NewImportsValidator(t *testing.T) {
	assertions := require.New(t)

	validator := NewImportsValidator()

	assertions.NotNil(validator)
}

func Test_ImportsValidator_WithBlacklist_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewImportsValidator()
	expectedBlacklist := []*regexp.Regexp{regexp.MustCompile("^abc$"), regexp.MustCompile("^def$")}

	h := handler.WithBlacklist([]string{"abc", "def"})

	assertions.Equal(handler, h)
	assertions.Equal(expectedBlacklist, h.blacklistImports)
}

func Test_ImportsValidator_WithWhitelist_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewImportsValidator()
	expectedWhitelist := []*regexp.Regexp{regexp.MustCompile("^abc$"), regexp.MustCompile("^def$")}

	h := handler.WithWhitelist([]string{"abc", "def"})

	assertions.Equal(handler, h)
	assertions.Equal(expectedWhitelist, h.whitelistImports)
}

func Test_ImportsValidator_WithFailFunc_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewImportsValidator()
	var failFunc validatorFailFunc
	failFunc = func() {}

	h := handler.WithFailFunc(failFunc)

	assertions.Equal(handler, h)
	assertions.Equal(fmt.Sprintf("%v", failFunc), fmt.Sprintf("%v", h.failFunc))
}

func Test_ImportsValidator_Validate_BlockedImport(t *testing.T) {
	assertions := require.New(t)
	failed := false
	NewImportsValidator().
		WithFailFunc(func() {
			failed = true
		}).
		WithBlacklist([]string{"fmt"}).
		Register()
	fset, node := parseSource(srcTest)

	inspector.Inspect(fset, node)

	assertions.True(failed)
}

func Test_ImportsValidator_Validate_UnknownImport(t *testing.T) {
	assertions := require.New(t)
	failed := false
	NewImportsValidator().
		WithFailFunc(func() {
			failed = true
		}).
		WithWhitelist([]string{"abc"}).
		Register()
	fset, node := parseSource(srcTest)

	inspector.Inspect(fset, node)

	assertions.True(failed)
}

func Test_ImportsValidator_Validate_KnownImport(t *testing.T) {
	assertions := require.New(t)
	failed := false
	NewImportsValidator().
		WithFailFunc(func() {
			failed = true
		}).
		WithWhitelist([]string{"fmt"}).
		Register()
	fset, node := parseSource(srcTest)

	inspector.Inspect(fset, node)

	assertions.False(failed)
}
