package renderer

import (
	"vpcpeeringviz/internal/node"
)

type Renderer interface {
	Render(nodes []node.Node) error
}
