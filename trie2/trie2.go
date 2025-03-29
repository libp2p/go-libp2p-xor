package trie2

type Key interface {
	Equal(Key) bool
	BitAt(int) byte
	Len() int
}

type Trie struct {
	Branch [2]*Trie
	Key    Key
}

func (trie *Trie) IsEmpty() bool {
	return trie.Key == nil
}

func (trie *Trie) IsLeaf() bool {
	return trie.Branch[0] == nil && trie.Branch[1] == nil
}

func (trie *Trie) IsEmptyLeaf() bool {
	return trie.IsEmpty() && trie.IsLeaf()
}

func (trie *Trie) IsNonEmptyLeaf() bool {
	return !trie.IsEmpty() && trie.IsLeaf()
}

func (trie *Trie) Size() int {
	return trie.SizeAtDepth(0)
}

func (trie *Trie) SizeAtDepth(depth int) int {
	if trie.IsLeaf() {
		if trie.IsEmpty() {
			return 0
		} else {
			return 1
		}
	} else {
		var here int
		if trie.IsEmpty() {
			here = 0
		} else {
			here = 1
		}
		return here + trie.Branch[0].SizeAtDepth(depth+1) + trie.Branch[1].SizeAtDepth(depth+1)
	}
}
