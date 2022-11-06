package graph

import (
	"vpcpeeringviz/internal/node"
)

type Grapher interface {
	Render(nodes []node.Node) error
}
