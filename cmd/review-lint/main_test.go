package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Log_Validate(t *testing.T) {
	assertions := require.New(t)
	file, err := ioutil.TempFile(".", "validate")
	if err != nil {
		panic(err)
	}
	file.WriteString(`
		package main

		import (
			"context"

			"github.com/mercadolibre/fury_go-core/pkg/log"
		)

		func NonCompliantLog() {
			log.Error(context.Background(), "non compliant string")
		}
	`)
	defer os.Remove(file.Name())
	args := []string{file.Name()}
	exitCode := 0
	exitFunc := func(code int) { exitCode = code }

	validate(args, exitFunc)

	assertions.Equal(1, exitCode)
}

func Test_Anything_Validate(t *testing.T) {
	assertions := require.New(t)
	file, err := ioutil.TempFile(".", "validate")
	if err != nil {
		panic(err)
	}
	file.WriteString(`
		package main

		import (
			"github.com/stretchr/testify/mock"
		)
		
		func NonCompliantMock() {
			m := &mock.Mock{}
			m.On("MyMethod", mock.Anything)
		}
	`)
	defer os.Remove(file.Name())
	args := []string{file.Name()}
	exitCode := 0
	exitFunc := func(code int) { exitCode = code }

	validate(args, exitFunc)

	assertions.Equal(1, exitCode)
}

func Test_Import_Validate(t *testing.T) {
	assertions := require.New(t)
	file, err := ioutil.TempFile(".", "validate")
	if err != nil {
		panic(err)
	}
	file.WriteString(`
		package main

		import (
			"log"
		)
		
		func NonCompliantImport() {
			log.Fatal("non compliant string")
		}
	`)
	defer os.Remove(file.Name())
	args := []string{file.Name()}
	exitCode := 0
	exitFunc := func(code int) { exitCode = code }

	validate(args, exitFunc)

	assertions.Equal(1, exitCode)
}

func Test_Fail(t *testing.T) {
	assertions := require.New(t)

	fail()

	assertions.Equal(1, exitCode)
}
