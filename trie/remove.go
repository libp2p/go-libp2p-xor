package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// Remove removes the key q from the trie. Remove mutates the trie.
// TODO: Also implement an immutable version of Add.
func (trie *Trie) Remove(q key.Key) (removedDepth int, removed bool) {
	return trie.RemoveAtDepth(0, q)
}

func (trie *Trie) RemoveAtDepth(depth int, q key.Key) (reachedDepth int, removed bool) {
	if qb := trie.Branch[q.BitAt(depth)]; qb != nil {
		if d, ok := qb.RemoveAtDepth(depth+1, q); ok {
			trie.shrink()
			return d, true
		} else {
			return d, false
		}
	} else {
		if trie.Key != nil && key.Equal(q, trie.Key) {
			trie.Key = nil
			return depth, true
		} else {
			return depth, false
		}
	}
}
