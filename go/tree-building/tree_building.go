// Go exercise tree-building.
package tree

import (
	"errors"
	"sort"
)

// A record in the linear representation of web forum comments.
type Record struct {
	ID     int
	Parent int
}

// A node in the tree representation of web forum comments.
type Node struct {
	ID       int
	Children []*Node
}

// Convert the linear representation of web forum comments into the tree representation.
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i int, j int) bool { return records[i].ID < records[j].ID })
	nodes := make(map[int]*Node)
	for id, record := range records {
		if id != record.ID {
			return nil, errors.New("Non-continuous record IDs")
		}
		nodes[record.ID] = &Node{ID: record.ID}
	}
	for _, record := range records {
		if record.ID == 0 {
			if record.Parent == 0 {
				continue
			} else {
				return nil, errors.New("Root node must not have parent")
			}
		} else {
			if record.Parent >= record.ID {
				return nil, errors.New("Parent ID must be < child ID")
			}
		}
		parent := nodes[record.Parent]
		if parent.Children == nil {
			parent.Children = []*Node{nodes[record.ID]}
		} else {
			parent.Children = append(parent.Children, nodes[record.ID])
		}
	}
	return nodes[0], nil
}
