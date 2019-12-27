package evaluator

import (
	"github.com/ottermad/gointerpreter/ast"
	"github.com/ottermad/gointerpreter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
