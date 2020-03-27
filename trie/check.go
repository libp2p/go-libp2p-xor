package trie

import (
	"github.com/libp2p/go-libp2p-xor/key"
)

// CheckInvariant panics of the trie does not meet its invariant.
func (trie *Trie) CheckInvariant() {
	trie.checkInvariant(0, nil)
}

func (trie *Trie) checkInvariant(depth int, pathSoFar *triePath) {
	switch {
	case trie.IsEmptyLeaf(): // ok
	case trie.IsNonEmptyLeaf():
		if !pathSoFar.matchesKey(trie.Key) {
			panic("key found at invalid location in trie")
		}
	default:
		if trie.IsEmpty() {
			b0, b1 := trie.Branch[0], trie.Branch[1]
			b0.checkInvariant(depth+1, pathSoFar.Push(0))
			b1.checkInvariant(depth+1, pathSoFar.Push(1))
			switch {
			case b0.IsEmptyLeaf() && b1.IsEmptyLeaf():
				panic("intermediate node with two empty leaves")
			case b0.IsEmptyLeaf() && b1.IsNonEmptyLeaf():
				panic("intermediate node with one empty leaf")
			case b0.IsNonEmptyLeaf() && b1.IsEmptyLeaf():
				panic("intermediate node with one empty leaf")
			}
		} else {
			panic("intermediate node with a key")
		}
	}
}

type triePath struct {
	parent *triePath
	bit    byte
}

func (p *triePath) Push(bit byte) *triePath {
	return &triePath{parent: p, bit: bit}
}

func (p *triePath) RootPath() []byte {
	if p == nil {
		return nil
	} else {
		return append(p.parent.RootPath(), p.bit)
	}
}

func (p *triePath) matchesKey(k key.Key) bool {
	// Slower, but more explicit:
	// for i, b := range p.RootPath() {
	// 	if k.BitAt(i) != b {
	// 		return false
	// 	}
	// }
	// return true
	ok, _ := p.walk(k, 0)
	return ok
}

func (p *triePath) walk(k key.Key, depthToLeaf int) (ok bool, depthToRoot int) {
	if p == nil {
		return true, 0
	} else {
		parOk, parDepthToRoot := p.parent.walk(k, depthToLeaf+1)
		return k.BitAt(parDepthToRoot) == p.bit && parOk, parDepthToRoot + 1
	}
}

func (p *triePath) String() string {
	return p.string(0)
}

func (p *triePath) string(depthToLeaf int) string {
	if p == nil {
		return ""
	} else {
		switch {
		case p.bit == 0:
			return p.parent.string(depthToLeaf+1) + "0"
		case p.bit == 1:
			return p.parent.string(depthToLeaf+1) + "1"
		default:
			panic("bit digit > 1")
		}
	}
}
