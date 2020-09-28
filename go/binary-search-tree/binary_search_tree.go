// Go exercise binary-search-tree
package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Create a new binary tree with one node containing the provided data.
func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

// Add data to the binary tree.
func (bst *SearchTreeData) Insert(data int) {
	switch {
	case data <= bst.data:
		if bst.left == nil {
			bst.left = Bst(data)
		} else {
			bst.left.Insert(data)
		}
	case data > bst.data:
		if bst.right == nil {
			bst.right = Bst(data)
		} else {
			bst.right.Insert(data)
		}
	}
}

// Do an in-order iteration of the value in the binary tree, and map each value to a string.
func (bst *SearchTreeData) MapString(mapper func(int) string) []string {
	data := bst.dataSlice()
	mapped := make([]string, len(data))
	for i, v := range data {
		mapped[i] = mapper(v)
	}
	return mapped
}

// Do an in-order iteration of the value in the binary tree, and map each value to an int.
func (bst *SearchTreeData) MapInt(mapper func(int) int) []int {
	data := bst.dataSlice()
	mapped := make([]int, len(data))
	for i, v := range data {
		mapped[i] = mapper(v)
	}
	return mapped
}

// Produce a sorted slice of integerers containing the sorted data in the binary tree.
func (bst *SearchTreeData) dataSlice() []int {
	slice := []int{}
	bst.recurseDataSlice(&slice)
	return slice
}

// Add the sorted values in the binary (sub-)tree to the end of the provided slice.
func (bst *SearchTreeData) recurseDataSlice(slice *[]int) {
	if bst.left != nil {
		bst.left.recurseDataSlice(slice)
	}
	*slice = append(*slice, bst.data)
	if bst.right != nil {
		bst.right.recurseDataSlice(slice)
	}
}
