package trie2

func (trie *Trie) Add(q Key) (insertedDepth int, insertedOK bool) {
	return trie.AddAtDepth(0, q)
}

func (trie *Trie) AddAtDepth(depth int, q Key) (insertedDepth int, insertedOK bool) {
	if q.Len() == depth {
		if trie.Key == nil {
			trie.Key = q
			return depth, true
		} else {
			if q.Equal(trie.Key) {
				return depth, false
			} else {
				// trie.Key.Len() > depth
				trie.Key, q = q, trie.Key
			}
		}
	}
	// q.Len() > depth
	switch {
	case trie.IsEmptyLeaf():
		trie.Key = q
		return depth, true
	case trie.IsNonEmptyLeaf():
		if q.Equal(trie.Key) {
			return depth, false
		} else {
			if trie.Key.Len() == depth {
				// both branches are nil
				trie.Branch[0], trie.Branch[1] = &Trie{}, &Trie{}
				return trie.Branch[q.BitAt(depth)].AddAtDepth(depth+1, q)
			} else { // trie.Key.Len() > depth
				p := trie.Key
				trie.Key = nil
				// both branches are nil
				trie.Branch[0], trie.Branch[1] = &Trie{}, &Trie{}
				trie.Branch[p.BitAt(depth)].Key = p
				return trie.Branch[q.BitAt(depth)].AddAtDepth(depth+1, q)
			}
		}
	default:
		return trie.Branch[q.BitAt(depth)].AddAtDepth(depth+1, q)
	}
}
