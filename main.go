package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"vpcpeeringviz/internal"
)

var (
	output string
	format string
)

func main() {
	outputFlag := flag.String("output", "peering.png", "image output path")
	formatFlag := flag.String("format", "png", "image format")
	flag.Parse()

	args := make(map[string]string)
	args["output"] = *outputFlag
	args["format"] = *formatFlag

	scanner := bufio.NewScanner(os.Stdin)

	err := app.Run(scanner, args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
