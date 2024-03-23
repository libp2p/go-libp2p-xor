package trie

import (
	"encoding/json"

	"github.com/libp2p/go-libp2p-xor/key"
)

// Trie is a trie for equal-length bit vectors, which stores values only in the leaves.
// Trie node invariants:
// (1) Either both branches are nil, or both are non-nil.
// (2) If branches are non-nil, key must be nil.
// (3) If both branches are leaves, then they are both non-empty (have keys).
type Trie[T any] struct {
	Branch [2]*Trie[T]
	Key    key.Key
	Data   T
}

func New[T any]() *Trie[T] {
	return &Trie[T]{}
}

func FromKeys[T any](k []key.Key) *Trie[T] {
	t := New[T]()
	for _, k := range k {
		t.Add(k)
	}
	return t
}

func FromKeysAtDepth[T any](depth int, k []key.Key) *Trie[T] {
	t := New[T]()
	for _, k := range k {
		t.AddAtDepth(depth, k)
	}
	return t
}

func (trie *Trie[T]) String() string {
	b, _ := json.Marshal(trie)
	return string(b)
}

func (trie *Trie[T]) Depth() int {
	return trie.DepthAtDepth(0)
}

func (trie *Trie[T]) DepthAtDepth(depth int) int {
	if trie.Branch[0] == nil && trie.Branch[1] == nil {
		return depth
	} else {
		return max(trie.Branch[0].DepthAtDepth(depth+1), trie.Branch[1].DepthAtDepth(depth+1))
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Size returns the number of keys added to the trie.
// In other words, it returns the number of non-empty leaves in the trie.
func (trie *Trie[T]) Size() int {
	return trie.SizeAtDepth(0)
}

func (trie *Trie[T]) SizeAtDepth(depth int) int {
	if trie.Branch[0] == nil && trie.Branch[1] == nil {
		if trie.IsEmpty() {
			return 0
		} else {
			return 1
		}
	} else {
		return trie.Branch[0].SizeAtDepth(depth+1) + trie.Branch[1].SizeAtDepth(depth+1)
	}
}

func (trie *Trie[T]) IsEmpty() bool {
	return trie.Key == nil
}

func (trie *Trie[T]) IsLeaf() bool {
	return trie.Branch[0] == nil && trie.Branch[1] == nil
}

func (trie *Trie[T]) IsEmptyLeaf() bool {
	return trie.IsEmpty() && trie.IsLeaf()
}

func (trie *Trie[T]) IsNonEmptyLeaf() bool {
	return !trie.IsEmpty() && trie.IsLeaf()
}

func (trie *Trie[T]) Copy() *Trie[T] {
	if trie.IsLeaf() {
		return &Trie[T]{Key: trie.Key}
	}

	return &Trie[T]{Branch: [2]*Trie[T]{
		trie.Branch[0].Copy(),
		trie.Branch[1].Copy(),
	}}
}

func (trie *Trie[T]) shrink() {
	b0, b1 := trie.Branch[0], trie.Branch[1]
	switch {
	case b0.IsEmptyLeaf() && b1.IsEmptyLeaf():
		trie.Branch[0], trie.Branch[1] = nil, nil
	case b0.IsEmptyLeaf() && b1.IsNonEmptyLeaf():
		trie.Key = b1.Key
		trie.Branch[0], trie.Branch[1] = nil, nil
	case b0.IsNonEmptyLeaf() && b1.IsEmptyLeaf():
		trie.Key = b0.Key
		trie.Branch[0], trie.Branch[1] = nil, nil
	}
}
