package validators

import (
	"fmt"
	"testing"

	"github.com/cmduquemeli/review-lint/inspector"
	"github.com/stretchr/testify/require"
)

func Test_NewLogValidator(t *testing.T) {
	assertions := require.New(t)

	validator := NewLogValidator()

	assertions.NotNil(validator)
}

func Test_LogValidator_WithPackageImport_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewLogValidator()
	expectedPackageImport := "packageImport"

	h := handler.WithPackageImport(expectedPackageImport)

	assertions.Equal(handler, h)
	assertions.Equal(expectedPackageImport, h.packageImport)
}

func Test_LogValidator_WithFunctionName_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewLogValidator()
	expectedFunctionName := "functionName"

	h := handler.WithFunctionName(expectedFunctionName)

	assertions.Equal(handler, h)
	assertions.Equal(expectedFunctionName, h.functionName)
}

func Test_LogValidator_WithFailFunc_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewLogValidator()
	var failFunc validatorFailFunc
	failFunc = func() {}

	h := handler.WithFailFunc(failFunc)

	assertions.Equal(handler, h)
	assertions.Equal(fmt.Sprintf("%v", failFunc), fmt.Sprintf("%v", h.failFunc))
}

func Test_LogValidator_Validate_ReservedWord(t *testing.T) {
	assertions := require.New(t)
	failed := false
	NewLogValidator().
		WithFailFunc(func() {
			failed = true
		}).
		Register()
	srcTest := `
		package test

		import (
			"context"
			"github.com/mercadolibre/fury_go-core/pkg/log"
		)

		func Test() {
			log.Error(context.Background(), "msg")
		}
	`
	fset, node := parseSource(srcTest)

	inspector.Inspect(fset, node)

	assertions.True(failed)
}

func Test_LogValidator_Validate_NonCompliantString(t *testing.T) {
	assertions := require.New(t)
	failed := false
	NewLogValidator().
		WithFailFunc(func() {
			failed = true
		}).
		Register()
	srcTest := `
		package test

		import (
			"context"
			"github.com/mercadolibre/fury_go-core/pkg/log"
		)

		func Test() {
			log.Error(context.Background(), "Hola Mundo")
		}
	`
	fset, node := parseSource(srcTest)

	inspector.Inspect(fset, node)

	assertions.True(failed)
}

func Test_LogValidator_Validate_OK(t *testing.T) {
	assertions := require.New(t)
	failed := false
	NewLogValidator().
		WithFailFunc(func() {
			failed = true
		}).
		Register()
	srcTest := `
		package test

		import (
			"context"
			"github.com/mercadolibre/fury_go-core/pkg/log"
		)

		func Test() {
			log.Error(context.Background(), "correct_message")
		}
	`
	fset, node := parseSource(srcTest)

	inspector.Inspect(fset, node)

	assertions.False(failed)
}
