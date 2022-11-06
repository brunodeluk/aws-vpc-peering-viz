package parser

import (
	"bufio"
	"strings"

	"vpcpeeringviz/internal/node"
)

func Parse(scanner *bufio.Scanner) []node.Node {
	nodes := make([]node.Node, 0)
	curNodesCount := -1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "VPCPEERINGCONNECTIONS") {
			peeringID := strings.Fields(line)[1]
			nodes = append(nodes, node.Node{
				PeeringID: peeringID,
			})
			curNodesCount += 1
		}

		if strings.HasPrefix(line, "ACCEPTERVPCINFO") {
			fields := strings.Fields(line)
			accepter := fields[len(fields)-1]
			nodes[curNodesCount].Accepter = accepter
		}

		if strings.HasPrefix(line, "REQUESTERVPCINFO") {
			fields := strings.Fields(line)
			requester := fields[len(fields)-1]
			nodes[curNodesCount].Requester = requester
		}
	}
	return nodes
}
