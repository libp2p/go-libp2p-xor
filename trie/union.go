package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

func UnionKeySlices(left, right []key.Key) []key.Key {
	result := append([]key.Key{}, left...)
	for _, item := range right {
		if !keyIsIn(item, result) {
			result = append(result, item)
		}
	}
	return result
}

func Union(left, right *Trie) *Trie {
	return UnionAtDepth(0, left, right)
}

func UnionAtDepth(depth int, left, right *Trie) *Trie {
	switch {
	case left.IsLeaf() && right.IsLeaf():
		switch {
		case left.IsEmpty() && right.IsEmpty():
			return &Trie{}
		case !left.IsEmpty() && right.IsEmpty():
			return &Trie{Key: left.Key}
		case left.IsEmpty() && !right.IsEmpty():
			return &Trie{Key: right.Key}
		case !left.IsEmpty() && !right.IsEmpty():
			return unionNonemptyLeaves(depth, left, right)
		}
	case !left.IsLeaf() && right.IsLeaf():
		return unionLeaf(depth, left, right)
	case left.IsLeaf() && !right.IsLeaf():
		return unionLeaf(depth, right, left)
	case !left.IsLeaf() && !right.IsLeaf():
		return &Trie{Branch: [2]*Trie{
			UnionAtDepth(depth+1, left.Branch[0], right.Branch[0]),
			UnionAtDepth(depth+1, left.Branch[1], right.Branch[1]),
		}}
	}
	panic("unreachable")
}

func unionLeaf(depth int, trie, leaf *Trie) *Trie {
	if leaf.IsEmpty() {
		return trie
	} else {
		dir := leaf.Key.BitAt(depth)
		copy := &Trie{}
		copy.Branch[dir] = UnionAtDepth(depth+1, trie.Branch[dir], leaf)
		copy.Branch[1-dir] = trie.Branch[1-dir]
		return copy
	}
}

func unionNonemptyLeaves(depth int, left, right *Trie) *Trie {
	switch {
	case left.Key.NormInt().Cmp(right.Key.NormInt()) == 0:
		return &Trie{Key: left.Key}
	case left.Key.BitAt(depth) == right.Key.BitAt(depth):
		return UnionAtDepth(depth, left.grow(depth), right.grow(depth))
	case left.Key.NormInt().Cmp(right.Key.NormInt()) < 0:
		return &Trie{Branch: [2]*Trie{
			&Trie{Key: left.Key},
			&Trie{Key: right.Key},
		}}
	case left.Key.NormInt().Cmp(right.Key.NormInt()) > 0:
		return &Trie{Branch: [2]*Trie{
			&Trie{Key: right.Key},
			&Trie{Key: left.Key},
		}}
	}
	panic("unreachable")
}
