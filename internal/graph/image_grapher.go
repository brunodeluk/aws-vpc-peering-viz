package graph

import (
	"vpcpeeringviz/internal/node"
)

type ImageGrapher struct {}

func (ImageGrapher) Render(nodes []node.Node) error {
	return nil
}
