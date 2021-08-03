package trie

import (
	"testing"

	"github.com/libp2p/go-libp2p-xor/key"
)

func TestInsertRemove(t *testing.T) {
	r := New()
	testSeq(r, t)
	testSeq(r, t)
}

func testSeq(r *Trie, t *testing.T) {
	for _, s := range testInsertSeq {
		depth, _ := r.Add(key.BytesKey(s.key))
		if depth != s.insertedDepth {
			t.Errorf("inserting expected depth %d, got %d", s.insertedDepth, depth)
		}
	}
	for _, s := range testRemoveSeq {
		depth, _ := r.Remove(key.BytesKey(s.key))
		if depth != s.reachedDepth {
			t.Errorf("removing expected depth %d, got %d", s.reachedDepth, depth)
		}
	}
}

func TestCopy(t *testing.T) {
	for _, sample := range testAddSamples {
		trie := FromKeys(sample.Keys)
		copy := trie.Copy()
		if d := copy.CheckInvariant(); d != nil {
			t.Fatalf("trie invariant discrepancy: %v", d)
		}
		if trie == copy {
			t.Errorf("Expected trie copy not to be the same reference as original")
		}
		if !Equal(trie, copy) {
			t.Errorf("Expected tries to be equal, original: %v\n, copy: %v\n", trie, copy)
		}
	}
}

var testInsertSeq = []struct {
	key           []byte
	insertedDepth int
}{
	{key: []byte{0x00}, insertedDepth: 0},
	{key: []byte{0x80}, insertedDepth: 1},
	{key: []byte{0x10}, insertedDepth: 4},
	{key: []byte{0xc0}, insertedDepth: 2},
	{key: []byte{0x20}, insertedDepth: 3},
}

var testRemoveSeq = []struct {
	key          []byte
	reachedDepth int
}{
	{key: []byte{0x00}, reachedDepth: 4},
	{key: []byte{0x10}, reachedDepth: 3},
	{key: []byte{0x20}, reachedDepth: 1},
	{key: []byte{0x80}, reachedDepth: 2},
	{key: []byte{0xc0}, reachedDepth: 0},
}
