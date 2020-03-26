package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// Find looks for the key q in the trie.
// It returns the depth of the leaf reached along the path of q, regardless of whether q was found in that leaf.
// It also returns a boolean flag indicating whether the key was found.
func (trie *Trie) Find(q key.Key) (reachedDepth int, found bool) {
	return trie.FindAtDepth(0, q)
}

func (trie *Trie) FindAtDepth(depth int, q key.Key) (reachedDepth int, found bool) {
	if qb := trie.Branch[q.BitAt(depth)]; qb != nil {
		return qb.FindAtDepth(depth+1, q)
	} else {
		if trie.Key == nil {
			return depth, false
		} else {
			return depth, key.Equal(trie.Key, q)
		}
	}
}
