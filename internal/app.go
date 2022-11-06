package app

import (
	"bufio"

	"vpcpeeringviz/internal/parser"
	"vpcpeeringviz/internal/renderer"
)

func Run(scanner *bufio.Scanner, args map[string]string) error {
	nodes := parser.Parse(scanner)

	config := render.Config{
		Output: args["output"],
		Layout:  args["format"]
	}

	err := renderer.Render(nodes, config)
	return err
}
