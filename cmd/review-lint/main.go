package main

import (
	"os"

	"github.com/cmduquemeli/review-lint/inspector"
	"github.com/cmduquemeli/review-lint/validators"
)

var exitCode = 0

func main() {
	validate(os.Args[1:], os.Exit)
}

func validate(args []string, exitFunc func(code int)) {
	validators.NewLogValidator().
		WithFailFunc(fail).
		Register()

	validators.NewMockAnythingValidator().
		WithFailFunc(fail).
		Register()

	validators.NewImportsValidator().
		WithFailFunc(fail).
		WithBlacklist(importsBlackList).
		WithWhitelist(importsWhiteList).
		Register()

	for _, filePath := range args {
		inspector.File(filePath)
	}

	exitFunc(exitCode)
}

func fail() {
	exitCode = 1
}
