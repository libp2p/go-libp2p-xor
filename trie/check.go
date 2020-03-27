package trie

// CheckInvariant panics of the trie does not meet its invariant.
func (trie *Trie) CheckInvariant() {
	switch {
	case trie.IsLeaf():
		return
	default:
		if trie.IsEmpty() {
			b0, b1 := trie.Branch[0], trie.Branch[1]
			b0.CheckInvariant()
			b1.CheckInvariant()
			switch {
			case b0.IsEmptyLeaf() && b1.IsEmptyLeaf():
				panic("intermediate node with two empty leaves")
			case b0.IsEmptyLeaf() && b1.IsNonEmptyLeaf():
				panic("intermediate node with one empty leaf")
			case b0.IsNonEmptyLeaf() && b1.IsEmptyLeaf():
				panic("intermediate node with one empty leaf")
			}
		} else {
			panic("intermediate node with a key")
		}
	}
}
