package kademlia

import (
	"sort"
	"testing"

	"github.com/libp2p/go-libp2p-xor/key"
	"github.com/libp2p/go-libp2p-xor/trie"
)

func randomTrie(count int, keySizeByte int) *trie.Trie {
	t := trie.New()
	for i := 0; i < count; i++ {
		t.Add(randomKey(keySizeByte))
	}
	return t
}

func TestClosestN(t *testing.T) {
	keySizeByte := 16
	root := randomTrie(100, keySizeByte)
	all := root.List()
	for count := 0; count <= 100; count += 10 {
		target := randomKey(keySizeByte)
		closest := ClosestN(target, root, count)
		if len(closest) != count {
			t.Fatalf("expected %d closest, found %d", count, len(closest))
		}
		sort.Slice(all, func(i, j int) bool {
			return string(key.Xor(all[i], target)) < string(key.Xor(all[j], target))
		})
		for i := range closest {
			if !key.Equal(closest[i], all[i]) {
				t.Errorf("wrong closest peer at offset %d: got %s, expected %s", i,
					closest[i], all[i])
			}
		}
	}
}

func closestTrivial(target key.Key, keys []key.Key, count int) []key.Key {
	sort.Slice(keys, func(i, j int) bool {
		return string(key.Xor(keys[i], target)) < string(key.Xor(keys[j], target))
	})
	return keys[:count]
}

// ensure the benchmark works.
var _x int

func BenchmarkClosestN(b *testing.B) {
	keySizeByte := 16
	root := randomTrie(100000, keySizeByte)
	count := 20
	target := randomKey(keySizeByte)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_x += len(ClosestN(target, root, count))
	}
}

func BenchmarkClosestTrivial(b *testing.B) {
	keySizeByte := 16
	root := randomTrie(100000, keySizeByte)
	keys := root.List()
	count := 20
	target := randomKey(keySizeByte)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_x += len(closestTrivial(target, keys, count))
	}
}
