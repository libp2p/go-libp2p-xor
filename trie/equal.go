package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

func Equal(p, q *XorTrie) bool {
	switch {
	case p.isLeaf() && q.isLeaf():
		return key.Equal(p.Key, q.Key)
	case !p.isLeaf() && !q.isLeaf():
		return Equal(p.Branch[0], q.Branch[0]) && Equal(p.Branch[1], q.Branch[1])
	}
	return false
}
