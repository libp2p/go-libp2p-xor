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
	left := FromKeys(sample.LeftKeys)
	right := FromKeys(sample.RightKeys)
	trie := Union(left, right)
	expected := FromKeys(UnionKeySlices(sample.LeftKeys, sample.RightKeys))

	if !Equal(trie, expected) {
		t.Errorf("union does not have expected values")
	}
	nodesMap := trieNodes(left, make(map[*Trie]bool))
	nodesMap = trieNodes(right, nodesMap)
	testTrieNotSameReference(t, nodesMap, trie)
}

func testTrieNotSameReference(t *testing.T, nodesMap map[*Trie]bool, union *Trie) {
	if union == nil {
		return
	}
	if nodesMap[union] {
		t.Errorf("Found a reference to original node in union: %v\n", union)
	}
	testTrieNotSameReference(t, nodesMap, union.Branch[0])
	testTrieNotSameReference(t, nodesMap, union.Branch[1])
}

func trieNodes(trie *Trie, nodesMap map[*Trie]bool) map[*Trie]bool {
	if trie == nil {
		return nodesMap
	}
	nodesMap[trie] = true
	nodesMap = trieNodes(trie.Branch[0], nodesMap)
	nodesMap = trieNodes(trie.Branch[1], nodesMap)
	return nodesMap
}
