package trie

import (
	"testing"

	"github.com/libp2p/go-libp2p-xor/key"
)

// Verify mutable and immutable add do the same thing.
func TestMutableAndImmutableAddSame(t *testing.T) {
	for _, s := range testAddSameSamples {
		mut := New()
		immut := New()
		for _, k := range s.Keys {
			mut.Add(k)
			immut = Add(immut, k)
		}
		if !Equal(mut, immut) {
			t.Errorf("mutable trie %v differs from immutable trie %v", mut, immut)
		}
	}
}

type testAddSameSample struct {
	Keys []key.Key
}

var testAddSameSamples = []*testAddSameSample{
	{Keys: []key.Key{{1, 3, 5, 7, 11, 13}}},
	{Keys: []key.Key{{11, 22, 23, 25, 27, 28, 31, 32, 33}}},
}
