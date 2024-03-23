package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// Add adds the key q to the trie. Add mutates the trie.
func (trie *Trie[T]) Add(q key.Key) (insertedDepth int, insertedOK bool) {
	return trie.AddAtDepth(0, q)
}

func (trie *Trie[T]) AddAtDepth(depth int, q key.Key) (insertedDepth int, insertedOK bool) {
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
			trie.Branch[0], trie.Branch[1] = &Trie[T]{}, &Trie[T]{}
			trie.Branch[p.BitAt(depth)].Key = p
			return trie.Branch[q.BitAt(depth)].AddAtDepth(depth+1, q)
		}
	default:
		return trie.Branch[q.BitAt(depth)].AddAtDepth(depth+1, q)
	}
}

// Add adds the key q to trie, returning a new trie.
// Add is immutable/non-destructive: The original trie remains unchanged.
func Add[T any](trie *Trie[T], q key.Key) *Trie[T] {
	return AddAtDepth(0, trie, q)
}

func AddAtDepth[T any](depth int, trie *Trie[T], q key.Key) *Trie[T] {
	switch {
	case trie.IsEmptyLeaf():
		return &Trie[T]{Key: q}
	case trie.IsNonEmptyLeaf():
		if key.Equal(trie.Key, q) {
			return trie
		} else {
			return trieForTwo[T](depth, trie.Key, q)
		}
	default:
		dir := q.BitAt(depth)
		s := &Trie[T]{}
		s.Branch[dir] = AddAtDepth(depth+1, trie.Branch[dir], q)
		s.Branch[1-dir] = trie.Branch[1-dir]
		return s
	}
}

func trieForTwo[T any](depth int, p, q key.Key) *Trie[T] {
	pDir, qDir := p.BitAt(depth), q.BitAt(depth)
	if qDir == pDir {
		s := &Trie[T]{}
		s.Branch[pDir] = trieForTwo[T](depth+1, p, q)
		s.Branch[1-pDir] = &Trie[T]{}
		return s
	} else {
		s := &Trie[T]{}
		s.Branch[pDir] = &Trie[T]{Key: p}
		s.Branch[qDir] = &Trie[T]{Key: q}
		return s
	}
}
