package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignOp struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	variable   node.Node
	expression node.Node
}
