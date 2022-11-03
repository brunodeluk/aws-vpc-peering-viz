package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/goccy/go-graphviz"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxBufSize := 128
	buff := make([]byte, maxBufSize)
	scanner.Buffer(buff, maxBufSize)

	var curPeering string
	var curAcc string
	var curReq string

	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "VPCPEERINGCONNECTIONS") {
			if curPeering != "" {
				a, err := graph.CreateNode(curReq)
				if err != nil {
					log.Fatal(err)
				}

				b, err := graph.CreateNode(curAcc)
				if err != nil {
					log.Fatal(err)
				}

				p, err := graph.CreateEdge(curPeering, a, b)
				if err != nil {
					log.Fatal(err)
				}

				p.SetLabel(curPeering)
			}

			curPeering = strings.Fields(line)[1]
		}

		if strings.Contains(line, "ACCEPTERVPCINFO") {
			accepterVpc := strings.Fields(line)[4]
			curAcc = accepterVpc
		}

		if strings.Contains(line, "REQUESTERVPCINFO") {
			requesterVpc := strings.Fields(line)[4]
			curReq = requesterVpc
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
		log.Fatal(err)
	}

	g.SetLayout(graphviz.CIRCO)

	if err := g.RenderFilename(graph, graphviz.PNG, "peerings.png"); err != nil {
		log.Fatal(err)
	}

}
