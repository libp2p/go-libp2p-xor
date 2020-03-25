package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// Intersect computes the intersection of the keys in p and q.
// p and q must be non-nil. The returned trie is never nil.
func Intersect(p, q *XorTrie) *XorTrie {
	return intersect(0, p, q)
}

func intersect(depth int, p, q *XorTrie) *XorTrie {
	switch {
	case p.isLeaf() && q.isLeaf():
		if p.isEmpty() || q.isEmpty() {
			return &XorTrie{} // empty set
		} else {
			if key.Equal(p.Key, q.Key) {
				return &XorTrie{Key: p.Key} // singleton
			} else {
				return &XorTrie{} // empty set
			}
		}
	case p.isLeaf() && !q.isLeaf():
		if p.isEmpty() {
			return &XorTrie{} // empty set
		} else {
			if _, found := q.find(depth, p.Key); found {
				return &XorTrie{Key: p.Key}
			} else {
				return &XorTrie{} // empty set
			}
		}
	case !p.isLeaf() && q.isLeaf():
		return Intersect(q, p)
	case !p.isLeaf() && !q.isLeaf():
		disjointUnion := &XorTrie{
			Branch: [2]*XorTrie{
				intersect(depth+1, p.Branch[0], q.Branch[0]),
				intersect(depth+1, p.Branch[1], q.Branch[1]),
			},
		}
		disjointUnion.shrink()
		return disjointUnion
	}
	panic("unreachable")
}
