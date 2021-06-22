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

	expected := FromKeys(UnionKeySlices(sample.LeftKeys, sample.RightKeys))
	actual := Union(left, right)
	if d := actual.CheckInvariant(); d != nil {
		t.Fatalf("trie invariant discrepancy: %v", d)
	}
	if !Equal(expected, actual) {
		t.Errorf("union of %v\n and %v\n expected %v\n got %v", left, right, expected, actual)
	}
}
