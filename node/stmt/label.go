package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Label struct {
	name       string
	attributes map[string]interface{}
	position *node.Position
	labelName  node.Node
}

func NewLabel(labelName node.Node) node.Node {
	return Label{
		"Label",
		map[string]interface{}{},
		nil,
		labelName,
	}
}

func (n Label) Name() string {
	return "Label"
}

func (n Label) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Label) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Label) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Label) Position() *node.Position {
	return n.position
}

func (n Label) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Label) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.labelName != nil {
		vv := v.GetChildrenVisitor("labelName")
		n.labelName.Walk(vv)
	}

	v.LeaveNode(n)
}
