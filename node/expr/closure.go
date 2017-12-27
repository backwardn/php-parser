package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Closure) Name() string {
	return "Closure"
}

type Closure struct {
	name        string
	params      []node.Node
	uses        []node.Node
	returnType  node.Node
	stmts       []node.Node
	isReturnRef bool
	isStatic    bool
}

func NewClosure(params []node.Node, uses []node.Node, returnType node.Node, stmts []node.Node, isStatic bool, isReturnRef bool) node.Node {
	return Closure{
		"Closure",
		params,
		uses,
		returnType,
		stmts,
		isReturnRef,
		isStatic,
	}
}

func (n Closure) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("isStatic", n.isStatic)
	v.Scalar("isReturnRef", n.isReturnRef)

	if n.params != nil {
		vv := v.Children("params")
		for _, nn := range n.params {
			nn.Walk(vv)
		}
	}

	if n.uses != nil {
		vv := v.Children("uses")
		for _, nn := range n.uses {
			nn.Walk(vv)
		}
	}

	if n.returnType != nil {
		vv := v.Children("returnType")
		n.returnType.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.Children("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}
}