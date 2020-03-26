package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// Add adds the key q to the trie. Add mutates the trie.
func (trie *Trie) Add(q key.Key) (insertedDepth int, insertedOK bool) {
	return trie.AddAtDepth(0, q)
}

func (trie *Trie) AddAtDepth(depth int, q key.Key) (insertedDepth int, insertedOK bool) {
	switch {
	case trie.IsEmptyLeaf():
		trie.Key = q
		return depth, true
	case trie.IsNonEmptyLeaf():
		if key.Equal(trie.Key, q) {
			return depth, false
		} else {
			p := trie.Key
			trie.Key = nil
			// both branches are nil
			trie.Branch[0], trie.Branch[1] = &Trie{}, &Trie{}
			trie.Branch[p.BitAt(depth)].AddAtDepth(depth+1, p)
			return trie.Branch[q.BitAt(depth)].AddAtDepth(depth+1, q)
		}
	default:
		return trie.Branch[q.BitAt(depth)].AddAtDepth(depth+1, q)
	}
}

// Add adds the key q to trie, returning a new trie.
// Add is immutable/non-destructive: The original trie remains unchanged.
func Add(trie *Trie, q key.Key) *Trie {
	return AddAtDepth(0, trie, q)
}

func AddAtDepth(depth int, trie *Trie, q key.Key) *Trie {
	switch {
	case trie.IsEmptyLeaf():
		return &Trie{Key: q}
	case trie.IsNonEmptyLeaf():
		if key.Equal(trie.Key, q) {
			return trie
		} else {
			dir := q.BitAt(depth)
			s := &Trie{}
			if q.BitAt(depth) == trie.Key.BitAt(depth) {
				s.Branch[dir] = AddAtDepth(depth+1, &Trie{Key: trie.Key}, q)
				s.Branch[1-dir] = &Trie{}
				return s
			} else {
				s.Branch[dir] = AddAtDepth(depth+1, &Trie{Key: trie.Key}, q)
				s.Branch[1-dir] = &Trie{}
			}
			return s
		}
	default:
		dir := q.BitAt(depth)
		s := &Trie{}
		s.Branch[dir] = AddAtDepth(depth+1, trie.Branch[dir], q)
		s.Branch[1-dir] = trie.Branch[1-dir]
		return s
	}
}
