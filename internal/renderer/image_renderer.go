package renderer

import (
	"vpcpeeringviz/internal/node"
)

type ImageRenderer struct {}

func (ImageRenderer) Render(nodes []node.Node) error {
	return nil
}
