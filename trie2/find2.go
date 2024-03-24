package trie2

// FindSubKeys finds all keys in the trie that are prefixes of the query key.
func (trie *Trie) FindSubKeys(q Key) (reachedDepth int, found []Key) {
	return trie.FindSubKeysAtDepth(0, q)
}

func (trie *Trie) FindSubKeysAtDepth(depth int, q Key) (reachedDepth int, found []Key) {
	if trie.IsLeaf() {
		if trie.IsEmpty() {
			return depth, nil
		} else {
			return depth, []Key{trie.Key}
		}
	} else {
		reachedDepth, found = trie.Branch[q.BitAt(depth)].FindSubKeysAtDepth(depth+1, q)
		if trie.IsEmpty() {
			return reachedDepth, found
		} else {
			return reachedDepth, append(found, trie.Key)
		}
	}
}

// FindSuperKeys finds all keys in the trie that are prefixed by the query key.
func (trie *Trie) FindSuperKeys(q Key) (reachedDepth int, found []Key) {
	return trie.FindSuperKeysAtDepth(0, q)
}

func (trie *Trie) FindSuperKeysAtDepth(depth int, q Key) (reachedDepth int, found []Key) {
	d, p := trie.Walk(depth, q)
	return d, p.List()
}

func (trie *Trie) Walk(depth int, q Key) (reachedDepth int, arrivedAt *Trie) {
	if trie.IsLeaf() {
		return depth, trie
	} else {
		return trie.Branch[q.BitAt(depth)].Walk(depth+1, q)
	}
}
