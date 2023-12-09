package day8

import (
	"fmt"
	"regexp"
	"strings"

	"AdityaHegde/adventofcode/go/utils"
)

type network struct {
	nodes    map[string]*node
	starts   []*node
	sequence string
}

type node struct {
	idx   int
	label string
	left  string
	right string
}

var nodeRegex = regexp.MustCompile(`^(\w*) = \((\w*), (\w*)\)$`)

func (n *network) parseNode(line string, idx int) {
	nodeMatches := nodeRegex.FindStringSubmatch(line)
	nd := &node{
		idx:   idx,
		label: nodeMatches[1],
		left:  nodeMatches[2],
		right: nodeMatches[3],
	}
	n.nodes[nodeMatches[1]] = nd
	if strings.HasSuffix(nodeMatches[1], "A") {
		n.starts = append(n.starts, nd)
	}
}

func parse(input []string) *network {
	n := &network{
		nodes:    make(map[string]*node),
		sequence: input[0],
	}

	for i, line := range input[2:] {
		n.parseNode(line, i)
	}

	return n
}

func partOne(input []string) int {
	n := parse(input)
	si := 0
	ns := "AAA"
	c := 0

	for ns != "ZZZ" {
		nd := n.nodes[ns]
		if n.sequence[si] == 'L' {
			ns = nd.left
		} else {
			ns = nd.right
		}
		si = (si + 1) % len(n.sequence)
		c++
	}

	return c
}

func partTwo(input []string) int64 {
	n := parse(input)
	cycles := make([]int64, len(n.starts))

	for i, nd := range n.starts {
		cy := utils.NewTwoDGrid[int]()
		c := 0
		si := 0

		for !cy.Has(si, nd.idx) {
			cy.Set(si, nd.idx, c)
			if n.sequence[si] == 'L' {
				nd = n.nodes[nd.left]
			} else {
				nd = n.nodes[nd.right]
			}
			si = (si + 1) % len(n.sequence)
			c++
		}

		fmt.Println(n.starts[i], c, cy.Grid[si][nd.idx])
		cycles[i] = int64(c - cy.Grid[si][nd.idx])
	}
	fmt.Println(cycles)

	return utils.LCM(cycles)
}

// 156549631109496832
// 15995167053923
