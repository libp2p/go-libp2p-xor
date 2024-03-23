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

func Union[T any](left, right *Trie[T]) *Trie[T] {
	return UnionAtDepth(0, left, right)
}

func UnionAtDepth[T any](depth int, left, right *Trie[T]) *Trie[T] {
	switch {
	case left.IsLeaf() && right.IsLeaf():
		switch {
		case left.IsEmpty() && right.IsEmpty():
			return &Trie[T]{}
		case !left.IsEmpty() && right.IsEmpty():
			return &Trie[T]{Key: left.Key}
		case left.IsEmpty() && !right.IsEmpty():
			return &Trie[T]{Key: right.Key}
		case !left.IsEmpty() && !right.IsEmpty():
			u := &Trie[T]{}
			u.AddAtDepth(depth, left.Key)
			u.AddAtDepth(depth, right.Key)
			return u
		}
	case !left.IsLeaf() && right.IsLeaf():
		return unionTrieAndLeaf(depth, left, right)
	case left.IsLeaf() && !right.IsLeaf():
		return unionTrieAndLeaf(depth, right, left)
	case !left.IsLeaf() && !right.IsLeaf():
		return &Trie[T]{Branch: [2]*Trie[T]{
			UnionAtDepth(depth+1, left.Branch[0], right.Branch[0]),
			UnionAtDepth(depth+1, left.Branch[1], right.Branch[1]),
		}}
	}
	panic("unreachable")
}

func unionTrieAndLeaf[T any](depth int, trie, leaf *Trie[T]) *Trie[T] {
	if leaf.IsEmpty() {
		return trie.Copy()
	} else {
		dir := leaf.Key.BitAt(depth)
		copy := &Trie[T]{}
		copy.Branch[dir] = UnionAtDepth(depth+1, trie.Branch[dir], leaf)
		copy.Branch[1-dir] = trie.Branch[1-dir].Copy()
		return copy
	}
}
