package trie2

// List returns a list of all keys in the trie.
func (trie *Trie) List() []Key {
	switch {
	case trie.IsEmptyLeaf():
		return nil
	case trie.IsNonEmptyLeaf():
		return []Key{trie.Key}
	case trie.IsEmpty():
		return append(trie.Branch[0].List(), trie.Branch[1].List()...)
	case !trie.IsEmpty():
		return append(
			[]Key{trie.Key},
			append(trie.Branch[0].List(), trie.Branch[1].List()...)...,
		)
	}
	panic("unreachable")
}
