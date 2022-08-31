package inspector

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type HandlerMock struct {
	mock.Mock
}

func (h *HandlerMock) Process(fset *token.FileSet, node ast.Node) {
	h.Called(fset, node)
}

func Test_RegisterHandler(t *testing.T) {
	assertions := require.New(t)
	handlers = map[string][]Handler{}
	nodeType := "nodeType"
	handler := HandlerMock{}

	RegisterHandler(nodeType, &handler)

	assertions.Equal(1, len(handlers[nodeType]))
}

func Test_RegisterTwoHandler(t *testing.T) {
	assertions := require.New(t)
	handlers = map[string][]Handler{}
	nodeType := "nodeType"
	handler1 := HandlerMock{}
	handler2 := HandlerMock{}

	RegisterHandler(nodeType, &handler1)
	RegisterHandler(nodeType, &handler2)

	assertions.Equal(2, len(handlers[nodeType]))
}
