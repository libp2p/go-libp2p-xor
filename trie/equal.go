package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

func Equal[T any](p, q *Trie[T]) bool {
	switch {
	case p.IsLeaf() && q.IsLeaf():
		return key.Equal(p.Key, q.Key)
	case !p.IsLeaf() && !q.IsLeaf():
		return Equal(p.Branch[0], q.Branch[0]) && Equal(p.Branch[1], q.Branch[1])
	}
	return false
}
