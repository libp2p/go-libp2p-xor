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
			u := &Trie{}
			u.AddAtDepth(depth, left.Key)
			u.AddAtDepth(depth, right.Key)
			return u
		}
	case !left.IsLeaf() && right.IsLeaf():
		return unionTrieAndLeaf(depth, left, right)
	case left.IsLeaf() && !right.IsLeaf():
		return unionTrieAndLeaf(depth, right, left)
	case !left.IsLeaf() && !right.IsLeaf():
		return &Trie{Branch: [2]*Trie{
			UnionAtDepth(depth+1, left.Branch[0], right.Branch[0]),
			UnionAtDepth(depth+1, left.Branch[1], right.Branch[1]),
		}}
	}
	panic("unreachable")
}

func unionTrieAndLeaf(depth int, trie, leaf *Trie) *Trie {
	if leaf.IsEmpty() {
		return &Trie{Branch: [2]*Trie{
			trie.Branch[0],
			trie.Branch[1],
		}}
	} else {
		dir := leaf.Key.BitAt(depth)
		copy := &Trie{}
		copy.Branch[dir] = UnionAtDepth(depth+1, trie.Branch[dir], leaf)
		copy.Branch[1-dir] = trie.Branch[1-dir]
		return copy
	}
}
