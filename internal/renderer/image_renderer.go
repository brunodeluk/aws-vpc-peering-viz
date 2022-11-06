package renderer

import (
	"vpcpeeringviz/internal/node"

	"github.com/goccy/go-graphviz"
)

type ImageRenderer struct {}

func (ImageRenderer) Render(nodes []node.Node, config Config) error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return err
	}
	defer func() {
		graph.Close()
		g.Close()
	}()

	for _, node := range nodes {
		a, err := graph.CreateNode(node.Accepter)
		if err != nil {
			return err
		}

		b, err := graph.CreateNode(node.Requester)
		if err != nil {
			return err
		}

		p, err := graph.CreateEdge(node.PeeringID, a, b)
		if err != nil {
			return err
		}

		p.SetLabel(node.PeeringID)
	}

	graphvizFormats := map[string]graphviz.Format{
		"png": graphviz.PNG,
		"svg": graphviz.SVG,
		"jpg": graphviz.JPG,
	}

	// var buf bytes.Buffer
	// if err := g.Render(graph, graphvizFormats[config.Format], &buf); err != nil {
	// 	return err
	// }

	g.SetLayout(graphviz.CIRCO)

	if err := g.RenderFilename(graph, graphvizFormats[config.Format], config.Output); err != nil {
		return err
	}

	return nil
}
