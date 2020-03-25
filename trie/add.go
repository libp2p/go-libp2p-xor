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
		s.branch[dir] = add(depth+1, trie.branch[dir], q)
		s.branch[1-dir] = trie.branch[1-dir]
		return s
	} else {
		if trie.key == nil {
			return &XorTrie{key: q}
		} else {
			if key.Equal(trie.key, q) {
				return trie
			} else {
				s := &XorTrie{}
				if q.BitAt(depth) == trie.key.BitAt(depth) {
					s.branch[dir] = add(depth+1, &XorTrie{key: trie.key}, q)
					s.branch[1-dir] = &XorTrie{}
					return s
				} else {
					s.branch[dir] = add(depth+1, &XorTrie{key: trie.key}, q)
					s.branch[1-dir] = &XorTrie{}
				}
				return s
			}
		}
	}
}
