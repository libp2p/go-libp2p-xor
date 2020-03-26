package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// Add adds the key q to the trie. Add mutates the trie.
// TODO: Also implement an immutable version of Add.
func (trie *Trie) Add(q key.Key) (insertedDepth int, insertedOK bool) {
	return trie.AddAtDepth(0, q)
}

func (trie *Trie) AddAtDepth(depth int, q key.Key) (insertedDepth int, insertedOK bool) {
	if qb := trie.Branch[q.BitAt(depth)]; qb != nil {
		return qb.AddAtDepth(depth+1, q)
	} else {
		if trie.Key == nil {
			trie.Key = q
			return depth, true
		} else {
			if key.Equal(trie.Key, q) {
				return depth, false
			} else {
				p := trie.Key
				trie.Key = nil
				// both Branches are nil
				trie.Branch[0], trie.Branch[1] = &Trie{}, &Trie{}
				trie.Branch[p.BitAt(depth)].AddAtDepth(depth+1, p)
				return trie.Branch[q.BitAt(depth)].AddAtDepth(depth+1, q)
			}
		}
	}
}

// Add adds the key q to trie, returning a new trie.
// Add is immutable/non-destructive: The original trie remains unchanged.
func Add(trie *Trie, q key.Key) *Trie {
	return AddAtDepth(0, trie, q)
}

func AddAtDepth(depth int, trie *Trie, q key.Key) *Trie {
	dir := q.BitAt(depth)
	if !trie.IsLeaf() {
		s := &Trie{}
		s.Branch[dir] = AddAtDepth(depth+1, trie.Branch[dir], q)
		s.Branch[1-dir] = trie.Branch[1-dir]
		return s
	} else {
		if trie.Key == nil {
			return &Trie{Key: q}
		} else {
			if key.Equal(trie.Key, q) {
				return trie
			} else {
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
		}
	}
}
