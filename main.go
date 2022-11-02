// 1. Read json input
// 2. Transform json input into graph data structure
// 3. Render the graph

package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Graph struct {
	M map[string][]string
}

func (g *Graph) Add(a string, b string) {
	if g.M[a] == nil {
		g.M[a] = make([]string, 1)
	}

	g.M[a] = append(g.M[a], b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxBufSize := 128
	buf := make([]byte, maxBufSize)
	scanner.Buffer(buf, maxBufSize)

	// currentPeering := Peering{}
	// peerings := make([]Graph, 0)
	var curPeering string
	var curAcc string
	var curReq string
	graph := Graph{
		M: make(map[string][]string),
	}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "VPCPEERINGCONNECTIONS") {
			if curPeering != "" {
				graph.Add(curReq, curAcc)
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

	log.Println(graph.M)

}
