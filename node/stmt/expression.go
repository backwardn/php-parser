package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Expression struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	expr       node.Node
}

func NewExpression(expr node.Node) node.Node {
	return Expression{
		"Expression",
		map[string]interface{}{},
		nil,
		expr,
	}
}

func (n Expression) Name() string {
	return "Expression"
}

func (n Expression) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Expression) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Expression) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Expression) Position() *node.Position {
	return n.position
}

func (n Expression) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Expression) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
