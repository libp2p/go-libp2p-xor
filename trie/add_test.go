package trie

import (
	"math/rand"
	"testing"

	"github.com/libp2p/go-libp2p-xor/key"
)

// Verify mutable and immutable add do the same thing.
func TestMutableAndImmutableAddSame(t *testing.T) {
	for _, s := range append(testAddSamples, randomTestAddSamples(100)...) {
		mut := New()
		immut := New()
		for _, k := range s.Keys {
			mut.Add(k)
			immut = Add(immut, k)
		}
		if d := mut.CheckInvariant(); d != nil {
			t.Fatalf("mutable trie invariant discrepancy: %v", d)
		}
		if d := immut.CheckInvariant(); d != nil {
			t.Fatalf("immutable trie invariant discrepancy: %v", d)
		}
		if !Equal(mut, immut) {
			t.Errorf("mutable trie %v differs from immutable trie %v", mut, immut)
		}
	}
}

func TestAddIsOrderIndependent(t *testing.T) {
	for _, s := range append(testAddSamples, randomTestAddSamples(100)...) {
		base := New()
		for _, k := range s.Keys {
			base.Add(k)
		}
		if d := base.CheckInvariant(); d != nil {
			t.Fatalf("base trie invariant discrepancy: %v", d)
		}
		for j := 0; j < 100; j++ {
			perm := rand.Perm(len(s.Keys))
			reordered := New()
			for i := range s.Keys {
				reordered.Add(s.Keys[perm[i]])
			}
			if d := reordered.CheckInvariant(); d != nil {
				t.Fatalf("reordered trie invariant discrepancy: %v", d)
			}
			if !Equal(base, reordered) {
				t.Errorf("trie %v differs from trie %v", base, reordered)
			}
		}
	}
}

type testAddSample struct {
	Keys []key.Key
}

var testAddSamples = []*testAddSample{
	{Keys: []key.Key{{1}, {3}, {5}, {7}, {11}, {13}}},
	{Keys: []key.Key{{11}, {22}, {23}, {25}, {27}, {28}, {31}, {32}, {33}}},
}

func randomTestAddSamples(count int) []*testAddSample {
	s := make([]*testAddSample, count)
	for i := range s {
		s[i] = randomTestAddSample(31, 2)
	}
	return s
}

func randomTestAddSample(setSize, keySizeByte int) *testAddSample {
	keySet := make([]key.Key, setSize)
	for i := range keySet {
		k := make(key.Key, keySizeByte)
		rand.Read(k)
		keySet[i] = k
	}
	return &testAddSample{
		Keys: keySet,
	}
}
