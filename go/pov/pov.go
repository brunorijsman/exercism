package pov

import "fmt"

// A node in the graph.
type Node struct {
	label     string  // Human-readable label for the node.
	toNodes   []*Node // Neighbor nodes for which there is an incoming arc into this node.
	fromNodes []*Node // Neighbor nodes for which there is an outgoing arc from this node.
}

// A graph.
type Graph struct {
	nodes map[string]*Node // The nodes in the graph. Each node keeps track of its neighbors.
}

// Construct a new graph without any nodes.
func New() *Graph {
	return &Graph{nodes: map[string]*Node{}}
}

// Add a node to the graph (it is not an error to add a node that already exists).
func (graph *Graph) AddNode(label string) {
	_, present := graph.nodes[label]
	if !present {
		graph.nodes[label] = &Node{label: label, toNodes: []*Node{}, fromNodes: []*Node{}}
	}
}

// Add an arc to the graph. If the nodes don't already exist, they will be added.
func (graph *Graph) AddArc(fromLabel, toLabel string) {
	graph.AddNode(fromLabel)
	graph.AddNode(toLabel)
	fromNode := graph.nodes[fromLabel]
	toNode := graph.nodes[toLabel]
	fromNode.toNodes = append(fromNode.toNodes, toNode)
	toNode.fromNodes = append(toNode.fromNodes, fromNode)
}

// Return a list of arcs represented as strings.
func (graph *Graph) ArcList() []string {
	result := []string{}
	for fromLabel, fromNode := range graph.nodes {
		for _, toNode := range fromNode.toNodes {
			result = append(result, fmt.Sprintf("%s -> %s", fromLabel, toNode.label))
		}
	}
	return result
}

// Return a new graph, re-rooted at newRoot.
func (graph *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	// The test program only stores trees in the graph, but allow for general graphs and handle
	// cycles in the graph correctly (hopefully... since there are no test cases for it)
	newGraph := New()
	root := graph.nodes[newRoot]
	visited := map[string]bool{}
	visitNode(graph, newGraph, nil, root, &visited)
	return newGraph
}

func visitNode(oldGraph, newGraph *Graph, prevNode, node *Node, visited *map[string]bool) {
	if (*visited)[node.label] {
		return
	}
	(*visited)[node.label] = true
	newGraph.AddNode(node.label)
	if prevNode != nil {
		newGraph.AddArc(prevNode.label, node.label)
	}
	for _, nextNode := range node.toNodes {
		visitNode(oldGraph, newGraph, node, nextNode, visited)
	}
	for _, nextNode := range node.fromNodes {
		visitNode(oldGraph, newGraph, node, nextNode, visited)
	}
}
