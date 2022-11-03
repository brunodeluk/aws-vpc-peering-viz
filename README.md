# AWS VPC Peering Visualizer (WIP)

A simple cmd tool to visualize your AWS VPC Peering topology

## Installation

1. Clone this repo
2. `go build .`
3. Optionally: `sudo cp vpc_graph /usr/local/bin/vpc_graph`

## How to use

```
aws vpc describe-vpc-peering-connections --output text | vpc_graph
```

The output is a `peerings.png` image located in your current dir path

