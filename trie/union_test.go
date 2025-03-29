package trie

import (
	"testing"
)

func TestUnion(t *testing.T) {
	for _, set := range testSetSamples {
		testUnion(t, set)
	}
}

func TestUnionRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		testUnion(t, randomTestSetSample(10, 21, 7))
	}
}

func TestUnionFromJSON(t *testing.T) {
	for _, json := range testJSONSamples {
		set := testSampleSetFromJSON(json)
		testUnion(t, set)
	}
}

func testUnion(t *testing.T, sample *testSetSample) {
	left := FromKeys[any](sample.LeftKeys)
	right := FromKeys[any](sample.RightKeys)
	trie := Union(left, right)
	expected := FromKeys[any](UnionKeySlices(sample.LeftKeys, sample.RightKeys))

	if !Equal(trie, expected) {
		t.Errorf("union does not have expected values")
	}
	nodesMap := trieNodes(left, make(map[*Trie[any]]bool))
	nodesMap = trieNodes(right, nodesMap)
	testTrieNotSameReference(t, nodesMap, trie)
}

func testTrieNotSameReference[T any](t *testing.T, nodesMap map[*Trie[T]]bool, union *Trie[T]) {
	if union == nil {
		return
	}
	if nodesMap[union] {
		t.Errorf("Found a reference to original node in union: %v\n", union)
	}
	testTrieNotSameReference(t, nodesMap, union.Branch[0])
	testTrieNotSameReference(t, nodesMap, union.Branch[1])
}

func trieNodes[T any](trie *Trie[T], nodesMap map[*Trie[T]]bool) map[*Trie[T]]bool {
	if trie == nil {
		return nodesMap
	}
	nodesMap[trie] = true
	nodesMap = trieNodes(trie.Branch[0], nodesMap)
	nodesMap = trieNodes(trie.Branch[1], nodesMap)
	return nodesMap
}
