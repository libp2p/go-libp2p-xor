package kademlia

import (
	"github.com/libp2p/go-libp2p-xor/key"
	"github.com/libp2p/go-libp2p-xor/trie"
)

// TableHealthReport describes the discrepancy between a node's routing table from the theoretical ideal,
// given knowledge of all nodes present in the network.
type TableHealthReport struct {
	// IdealDepth is the depth that the node's rouing table should have.
	IdealDepth int
	// ActualDepth is the depth that the node's routing table has.
	ActualDepth int
	// Bucket...
	Bucket []*BucketHealthReport
}

// BucketHealth describes the discrepancy between a node's routing bucket and the theoretical ideal,
// given knowledge of all nodes present in the network (aka the "known" nodes).
type BucketHealthReport struct {
	// MaxKnownContacts is the number of all known network nodes,
	// which are eligible to be in this bucket.
	MaxKnownContacts int
	// ActualKnownContacts is the number of known network nodes,
	// that are actually in the node's routing table.
	ActualKnownContacts int
	// ActualUnknownContacts is the number of contacts in the node's routing table,
	// that are not known to be in the network currently.
	ActualUnknownContacts int
}

// TableHealth computes the health report for a node,
// given its routing contacts and a list of all known nodes in the network currently.
func TableHealth(node key.Key, nodeContacts []key.Key, knownNodes *trie.Trie) *TableHealthReport {
	// Reconstruct the node's routing table as a trie
	nodeTable := trie.New()
	nodeTable.Add(node)
	for _, u := range nodeContacts {
		nodeTable.Add(u)
	}
	// Compute health report
	idealDepth, _ := knownNodes.Find(node)
	actualDepth, _ := nodeTable.Find(node)
	return &TableHealthReport{
		IdealDepth:  idealDepth,
		ActualDepth: actualDepth,
		Bucket:      BucketHealth(node, nodeTable, knownNodes),
	}
}

// BucketHealth computes the health report for each bucket in a node's routing table,
// given the node's routing table and a list of all known nodes in the network currently.
func BucketHealth(node key.Key, nodeTable, knownNodes *trie.Trie) []*BucketHealthReport {
	panic("u")
	// actualDepth, _ := nodeTable.Find(node)
	// bucket := makeXXX
	// return bucket
}
