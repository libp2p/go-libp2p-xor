package kademlia

import (
	"github.com/libp2p/go-libp2p-xor/key"
	"github.com/libp2p/go-libp2p-xor/trie"
)

// BucketAtDepth returns the bucket in the routing table at a given depth.
// A bucket at depth D holds contacts that share a prefix of exactly D bits with node.
func BucketAtDepth(node key.Key, table *trie.Trie, depth int) *trie.Trie {
	dir := node.BitAt(depth)
	if table.IsLeaf() {
		return nil
	} else {
		if depth == 0 {
			return table.Branch[1-dir]
		} else {
			return BucketAtDepth(node, table.Branch[dir], depth-1)
		}
	}
}
