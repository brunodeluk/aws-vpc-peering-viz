package renderer

import (
	"vpcpeeringviz/internal/node"
)

type Renderer interface {
	Render(nodes []node.Node) error
}

type Config struct {
	Output string
	Format string
	Layout string
}
