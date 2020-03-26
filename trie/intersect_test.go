package trie

import (
	"math/rand"
	"testing"

	"github.com/libp2p/go-libp2p-xor/key"
)

func TestIntersectRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		testIntersect(t, randomTestIntersectSample(10, 10, 5))
	}
}

func TestIntersect(t *testing.T) {
	for _, s := range testIntersectSamples {
		testIntersect(t, s)
	}
}

func testIntersect(t *testing.T, sample *testIntersectSample) {
	left, right, expected := New(), New(), New()
	for _, l := range sample.LeftKeys {
		left.Add(l)
	}
	for _, r := range sample.RightKeys {
		right.Add(r)
	}
	for _, s := range setIntersect(sample.LeftKeys, sample.RightKeys) {
		expected.Add(s)
	}
	got := Intersect(left, right)
	if !Equal(expected, got) {
		t.Errorf("intersection of %v and %v: expected %v, got %v",
			sample.LeftKeys, sample.RightKeys, expected, got)
	}
}

func setIntersect(left, right []key.Key) []key.Key {
	intersection := []key.Key{}
	for _, l := range left {
		for _, r := range right {
			if key.Equal(l, r) {
				intersection = append(intersection, r)
			}
		}
	}
	return intersection
}

func randomTestIntersectSample(leftSize, rightSize, intersectSize int) *testIntersectSample {
	keys := make([]key.Key, leftSize+rightSize-intersectSize)
	for i := range keys {
		keys[i] = key.Key{byte(rand.Intn(256))}
	}
	return &testIntersectSample{
		LeftKeys:  keys[:leftSize],
		RightKeys: keys[leftSize-intersectSize:],
	}
}

type testIntersectSample struct {
	LeftKeys  []key.Key
	RightKeys []key.Key
}

var testIntersectSamples = []*testIntersectSample{
	{
		LeftKeys:  []key.Key{{1, 2, 3}},
		RightKeys: []key.Key{{1, 3, 5}},
	},
	{
		LeftKeys:  []key.Key{{1, 2, 3, 4, 5, 6}},
		RightKeys: []key.Key{{3, 5, 7}},
	},
	{
		LeftKeys:  []key.Key{{23, 3, 7, 13, 17}},
		RightKeys: []key.Key{{2, 11, 17, 19, 23}},
	},
}
