package validators

import (
	"fmt"
	"testing"

	"github.com/cmduquemeli/review-lint/inspector"
	"github.com/stretchr/testify/require"
)

func Test_NewMockAnythingValidator(t *testing.T) {
	assertions := require.New(t)

	validator := NewMockAnythingValidator()

	assertions.NotNil(validator)
}

func Test_MockAnythingValidator_WithPackageImport_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewMockAnythingValidator()
	expectedPackageImport := "packageImport"

	h := handler.WithPackageImport(expectedPackageImport)

	assertions.Equal(handler, h)
	assertions.Equal(expectedPackageImport, h.packageImport)
}

func Test_MockAnythingValidator_WithFailFunc_Asign(t *testing.T) {
	assertions := require.New(t)
	handler := NewMockAnythingValidator()
	var failFunc validatorFailFunc
	failFunc = func() {}

	h := handler.WithFailFunc(failFunc)

	assertions.Equal(handler, h)
	assertions.Equal(fmt.Sprintf("%v", failFunc), fmt.Sprintf("%v", h.failFunc))
}

func Test_MockAnythingValidator_Validate_Anything(t *testing.T) {
	assertions := require.New(t)
	failed := false
	NewMockAnythingValidator().
		WithFailFunc(func() {
			failed = true
		}).
		Register()
	srcTest := `
		package test

		import (
			"context"
			"github.com/stretchr/testify/mock"
			"github.com/stretchr/testify/require"
		)

		func Test() {
			dbRepository := mock.Mock{}
			dbRepository.On("CountStatusDetail", mock.Anything, mock.Anything).Return(int64(5), nil)
		}
	`
	fset, node := parseSource(srcTest)

	inspector.Inspect(fset, node)

	assertions.True(failed)
}

func Test_MockAnythingValidator_Validate(t *testing.T) {
	assertions := require.New(t)
	failed := false
	NewMockAnythingValidator().
		WithFailFunc(func() {
			failed = true
		}).
		Register()
	srcTest := `
		package test

		import (
			"context"
			"github.com/stretchr/testify/mock"
			"github.com/stretchr/testify/require"
		)

		func Test() {
			dbRepository := mock.Mock{}
			dbRepository.On("CountStatusDetail", "arg1", "arg2").Return(int64(5), nil)
		}
	`
	fset, node := parseSource(srcTest)

	inspector.Inspect(fset, node)

	assertions.False(failed)
}
