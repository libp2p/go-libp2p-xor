package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// XorTrie is a trie for equal-length bit vectors, which stores values only in the leaves.
// XorTrie node invariants:
// (1) Either both branches are nil, or both are non-nil.
// (2) If branches are non-nil, key must be nil.
// (3) If both branches are leaves, then they are both non-empty (have keys).
type XorTrie struct {
	Branch [2]*XorTrie
	Key    key.Key
}

func New() *XorTrie {
	return &XorTrie{}
}

func (trie *XorTrie) Depth() int {
	return trie.depth(0)
}

func (trie *XorTrie) depth(depth int) int {
	if trie.Branch[0] == nil && trie.Branch[1] == nil {
		return depth
	} else {
		return max(trie.Branch[0].depth(depth+1), trie.Branch[1].depth(depth+1))
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (trie *XorTrie) Find(q key.Key) (reachedDepth int, found bool) {
	return trie.find(0, q)
}

func (trie *XorTrie) find(depth int, q key.Key) (reachedDepth int, found bool) {
	if qb := trie.Branch[q.BitAt(depth)]; qb != nil {
		return qb.find(depth+1, q)
	} else {
		if trie.Key == nil {
			return depth, false
		} else {
			return depth, key.Equal(trie.Key, q)
		}
	}
}

// Add adds the key q to the trie. Add mutates the trie.
// TODO: Also implement an immutable version of Add.
func (trie *XorTrie) Add(q key.Key) (insertedDepth int, insertedOK bool) {
	return trie.add(0, q)
}

func (trie *XorTrie) add(depth int, q key.Key) (insertedDepth int, insertedOK bool) {
	if qb := trie.Branch[q.BitAt(depth)]; qb != nil {
		return qb.add(depth+1, q)
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
				trie.Branch[0], trie.Branch[1] = &XorTrie{}, &XorTrie{}
				trie.Branch[p.BitAt(depth)].add(depth+1, p)
				return trie.Branch[q.BitAt(depth)].add(depth+1, q)
			}
		}
	}
}

// Remove removes the key q from the trie. Remove mutates the trie.
// TODO: Also implement an immutable version of Add.
func (trie *XorTrie) Remove(q key.Key) (removedDepth int, removed bool) {
	return trie.remove(0, q)
}

func (trie *XorTrie) remove(depth int, q key.Key) (reachedDepth int, removed bool) {
	if qb := trie.Branch[q.BitAt(depth)]; qb != nil {
		if d, ok := qb.remove(depth+1, q); ok {
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

func (trie *XorTrie) isEmpty() bool {
	return trie.Key == nil
}

func (trie *XorTrie) isLeaf() bool {
	return trie.Branch[0] == nil && trie.Branch[1] == nil
}

func (trie *XorTrie) isEmptyLeaf() bool {
	return trie.isEmpty() && trie.isLeaf()
}

func (trie *XorTrie) isNonEmptyLeaf() bool {
	return !trie.isEmpty() && trie.isLeaf()
}

func (trie *XorTrie) shrink() {
	b0, b1 := trie.Branch[0], trie.Branch[1]
	switch {
	case b0.isEmptyLeaf() && b1.isEmptyLeaf():
		trie.Branch[0], trie.Branch[1] = nil, nil
	case b0.isEmptyLeaf() && b1.isNonEmptyLeaf():
		trie.Key = b1.Key
		trie.Branch[0], trie.Branch[1] = nil, nil
	case b0.isNonEmptyLeaf() && b1.isEmptyLeaf():
		trie.Key = b0.Key
		trie.Branch[0], trie.Branch[1] = nil, nil
	}
}
