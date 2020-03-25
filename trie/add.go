package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// Add adds the key q to trie, returning a new trie.
// Add is immutable/non-destructive: The original trie remains unchanged.
func Add(trie *XorTrie, q key.Key) *XorTrie {
	return add(0, trie, q)
}

func add(depth int, trie *XorTrie, q key.Key) *XorTrie {
	dir := q.BitAt(depth)
	if !trie.isLeaf() {
		s := &XorTrie{}
		s.Branch[dir] = add(depth+1, trie.Branch[dir], q)
		s.Branch[1-dir] = trie.Branch[1-dir]
		return s
	} else {
		if trie.Key == nil {
			return &XorTrie{Key: q}
		} else {
			if key.Equal(trie.Key, q) {
				return trie
			} else {
				s := &XorTrie{}
				if q.BitAt(depth) == trie.Key.BitAt(depth) {
					s.Branch[dir] = add(depth+1, &XorTrie{Key: trie.Key}, q)
					s.Branch[1-dir] = &XorTrie{}
					return s
				} else {
					s.Branch[dir] = add(depth+1, &XorTrie{Key: trie.Key}, q)
					s.Branch[1-dir] = &XorTrie{}
				}
				return s
			}
		}
	}
}
