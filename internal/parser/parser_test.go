package parser_test

import (
	"bufio"
	"strings"
	"testing"
	"vpcpeeringviz/internal/parser"
)

func setUpScanner(input string) *bufio.Scanner {
	reader := strings.NewReader(input)
	return bufio.NewScanner(reader)
}

func TestParserWithNoInput_ReturnsEmptySlice(t *testing.T) {
	scanner := setUpScanner("")
	nodes := parser.Parse(scanner)
	if len(nodes) != 0 {
		t.Error("expected nodes slice to be empty")
	}
}

func TestInvalidInput_ReturnsEmptySlice(t *testing.T) {
	scanner := setUpScanner("this is a test line")
	nodes := parser.Parse(scanner)
	if len(nodes) != 0 {
		t.Error("expected nodes slice to be empty")
	}
}

func TestIncompleteInput_ReturnsEmptySlice(t *testing.T) {
	scanner := setUpScanner(`
	VPCPEERINGCONNECTIONS	pcx-0bc744e9e0469cdb0
	`)
	nodes := parser.Parse(scanner)
	if len(nodes) != 0 {
		t.Error("expected nodes slice to be empty")
	}
}

func TestInputWithOneVPCPeering_ReturnsOneNode(t *testing.T) {
	scanner := setUpScanner(`
VPCPEERINGCONNECTIONS	pcx-0bc744e9e0469cdb0
ACCEPTERVPCINFO	10.333.0.0/16	914808247515	us-east-1	vpc-0498a19dfaa5e25a7
REQUESTERVPCINFO	10.999.0.0/16	914808247515	us-east-1	vpc-06871633f00e71662
	`)

	nodes := parser.Parse(scanner)
	if len(nodes) != 1 {
		t.Errorf("expected 1 node; got %d\n", len(nodes))
	}

	node := nodes[0]
	if node.PeeringID != "pcx-0bc744e9e0469cdb0" {
		t.Errorf("expected pcx-0bc744e9e0469cdb0 peering id; got %s\n", node.PeeringID)
	}

	if node.Accepter != "vpc-0498a19dfaa5e25a7" {
		t.Errorf("expected vpc-0498a19dfaa5e25a7 accepter; got %s\n", node.Accepter)
	}

	if node.Requester != "vpc-06871633f00e71662" {
		t.Errorf("expected vpc-06871633f00e71662 requester; got %s\n", node.Requester)
	}
}

func TestValidInputWithIncompleteFields_ReturnsOneNode(t *testing.T) {
	scanner := setUpScanner(`
VPCPEERINGCONNECTIONS	pcx-0bc744e9e0469cdb0
ACCEPTERVPCINFO	10.333.0.0/16  vpc-0498a19dfaa5e25a7
REQUESTERVPCINFO	10.999.0.0/16	914808247515  vpc-06871633f00e71662
	`)

	nodes := parser.Parse(scanner)
	if len(nodes) != 1 {
		t.Errorf("expected 1 node; got %d\n", len(nodes))
	}

	node := nodes[0]
	if node.PeeringID != "pcx-0bc744e9e0469cdb0" {
		t.Errorf("expected pcx-0bc744e9e0469cdb0 peering id; got %s\n", node.PeeringID)
	}

	if node.Accepter != "vpc-0498a19dfaa5e25a7" {
		t.Errorf("expected vpc-0498a19dfaa5e25a7 accepter; got %s\n", node.Accepter)
	}

	if node.Requester != "vpc-06871633f00e71662" {
		t.Errorf("expected vpc-06871633f00e71662 requester; got %s\n", node.Requester)
	}
}

func TestFullValidInput_ReturnsOneNode(t *testing.T) {
	scanner := setUpScanner(`
VPCPEERINGCONNECTIONS	pcx-0bc744e9e0469cdb0
ACCEPTERVPCINFO	10.333.0.0/16	914808247515	us-east-1	vpc-0498a19dfaa5e25a7
CIDRBLOCKSET	10.333.0.0/16
PEERINGOPTIONS	True	False	False
REQUESTERVPCINFO	10.999.0.0/16	914808247515	us-east-1	vpc-06871633f00e71662
CIDRBLOCKSET	10.999.0.0/16
PEERINGOPTIONS	True	False	False
STATUS	active	Active
	`)

	nodes := parser.Parse(scanner)
	if len(nodes) != 1 {
		t.Errorf("expected 1 node; got %d\n", len(nodes))
	}

	node := nodes[0]
	if node.PeeringID != "pcx-0bc744e9e0469cdb0" {
		t.Errorf("expected pcx-0bc744e9e0469cdb0 peering id; got %s\n", node.PeeringID)
	}

	if node.Accepter != "vpc-0498a19dfaa5e25a7" {
		t.Errorf("expected vpc-0498a19dfaa5e25a7 accepter; got %s\n", node.Accepter)
	}

	if node.Requester != "vpc-06871633f00e71662" {
		t.Errorf("expected vpc-06871633f00e71662 requester; got %s\n", node.Requester)
	}
}
