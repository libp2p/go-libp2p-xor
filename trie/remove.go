package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// Remove removes the key q from the trie. Remove mutates the trie.
// TODO: Also implement an immutable version of Remove.
func (trie *Trie) Remove(q key.Key) (removedDepth int, removed bool) {
	return trie.RemoveAtDepth(0, q)
}

func (trie *Trie) RemoveAtDepth(depth int, q key.Key) (reachedDepth int, removed bool) {
	switch {
	case trie.IsEmptyLeaf():
		return depth, false
	case trie.IsNonEmptyLeaf():
		trie.Key = nil
		return depth, true
	default:
		if d, removed := trie.Branch[q.BitAt(depth)].RemoveAtDepth(depth+1, q); removed {
			trie.shrink()
			return d, true
		} else {
			return d, false
		}
	}
}
