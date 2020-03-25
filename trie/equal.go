package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

func Equal(p, q *XorTrie) bool {
	switch {
	case p.isLeaf() && q.isLeaf():
		return key.Equal(p.key, q.key)
	case !p.isLeaf() && !q.isLeaf():
		return Equal(p.branch[0], q.branch[0]) && Equal(p.branch[1], q.branch[1])
	}
	return false
}
