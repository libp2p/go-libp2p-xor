package trie

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/libp2p/go-libp2p-xor/key"
)

func TestIntersectRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		testIntersect(t, randomTestIntersectSample(10, 21, 7))
	}
}

func TestIntersect(t *testing.T) {
	for _, s := range testIntersectSamples {
		testIntersect(t, s)
	}
}

func TestIntersectFromJSON(t *testing.T) {
	for _, json := range testIntersectJSONSamples {
		s := testIntersectSampleFromJSON(json)
		testIntersect(t, s)
	}
}

func testIntersect(t *testing.T, sample *testIntersectSample) {
	left, right, expected := New(), New(), New()
	for _, l := range sample.LeftKeys {
		left.Add(l)
	}
	if d := left.CheckInvariant(); d != nil {
		t.Fatalf("left trie invariant discrepancy: %v", d)
	}
	for _, r := range sample.RightKeys {
		right.Add(r)
	}
	if d := right.CheckInvariant(); d != nil {
		t.Fatalf("right trie invariant discrepancy: %v", d)
	}
	for _, s := range setIntersect(sample.LeftKeys, sample.RightKeys) {
		expected.Add(s)
	}
	got := Intersect(left, right)
	if d := got.CheckInvariant(); d != nil {
		t.Fatalf("right trie invariant discrepancy: %v", d)
	}
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
		keys[i] = key.ByteKey(byte(rand.Intn(256)))
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

func testIntersectSampleFromJSON(srcJSON string) *testIntersectSample {
	s := &testIntersectSample{}
	if err := json.Unmarshal([]byte(srcJSON), s); err != nil {
		panic(err)
	}
	return s
}

var testIntersectSamples = []*testIntersectSample{
	{
		LeftKeys:  []key.Key{key.ByteKey(1), key.ByteKey(2), key.ByteKey(3)},
		RightKeys: []key.Key{key.ByteKey(1), key.ByteKey(3), key.ByteKey(5)},
	},
	{
		LeftKeys:  []key.Key{key.ByteKey(1), key.ByteKey(2), key.ByteKey(3), key.ByteKey(4), key.ByteKey(5), key.ByteKey(6)},
		RightKeys: []key.Key{key.ByteKey(3), key.ByteKey(5), key.ByteKey(7)},
	},
	{
		LeftKeys:  []key.Key{key.ByteKey(23), key.ByteKey(3), key.ByteKey(7), key.ByteKey(13), key.ByteKey(17)},
		RightKeys: []key.Key{key.ByteKey(2), key.ByteKey(11), key.ByteKey(17), key.ByteKey(19), key.ByteKey(23)},
	},
}

var testIntersectJSONSamples = []string{
	`
{
    "LeftKeys": [
        "gAlMTjoCy6ZDZFN/0okF65fxscCLVxnQhlJsyfp6uWU=",
        "IOZLVCi+OWdUqu0N7DX9T6sweK6RffGnBChy3I3G424=",
        "yIsHasiWkbMESfShoZ5yvS4fv6m0KnBSV6emNyMgYrg=",
        "qAISv6yZjs3WWlDC89iJUSdq45F0D/1y9fLnsPvavdA=",
        "aOl3hzLR2jArVpbEWaLOjH4QqoUo/7pJlmbY1mCHgYI=",
        "6P28fa97aZOOImD9TGO+x2+XJbzPSEh5rP6NqKoDj88=",
        "2Lh4bHSBoSvqloX+v7IYMvmq0k0bUcYAkQesOUPxLVo=",
        "eLqLdLjusZJlqDZT94U0PC4bdVQeUNaGwkOn9i5KRoI="
    ],
    "RightKeys": [
        "gAlMTjoCy6ZDZFN/0okF65fxscCLVxnQhlJsyfp6uWU=",
        "qAISv6yZjs3WWlDC89iJUSdq45F0D/1y9fLnsPvavdA="
    ]
}
	`,
	`
{
    "LeftKeys": [
        "BXmU8txOqn8ExHzXuXRtHm2XM99uD8lsgPo8OdDcNYE=",
        "hVzcamVYfIqs4IlrVIM1qRalqTh8OMrlAeqwgJTI2xo=",
        "JbILHEM3RcA+ksq0BvU+9Zfc+jnpsxPUQLe9lrHqBwc=",
        "VUeFiK5V64F8G2rvnIyoopfOzICF0h79FmeiLQqrVAI=",
        "tYRdsKlUTbTXOpgVjUZtzh2DRG0e5nPXIrkN60PI5GE=",
        "LatyclJiSPEaCoLxbabddv7Rqrsy+J1hf2Pd9BmmN1U=",
        "XX0wXrGF4IytkKmStxesXOiGFK+dm5ran6lWu7xNhIw=",
        "PcZb8TBEHhEqpFfaRWyhit3Uc03895uOkMgiiBgW9Uk="
    ],
    "RightKeys": [
        "BXmU8txOqn8ExHzXuXRtHm2XM99uD8lsgPo8OdDcNYE=",
        "hVzcamVYfIqs4IlrVIM1qRalqTh8OMrlAeqwgJTI2xo=",
        "VUeFiK5V64F8G2rvnIyoopfOzICF0h79FmeiLQqrVAI=",
        "tYRdsKlUTbTXOpgVjUZtzh2DRG0e5nPXIrkN60PI5GE=",
        "XX0wXrGF4IytkKmStxesXOiGFK+dm5ran6lWu7xNhIw="
    ]
}
	`,
}

func TestIntersectTriesFromJSON(t *testing.T) {
	for _, json := range testIntersectJSONTries {
		s := testIntersectTrieFromJSON(json)
		testIntersectTries(t, s)
	}
}

func testIntersectTries(t *testing.T, sample *testIntersectTrie) {
	if d := sample.LeftTrie.CheckInvariant(); d != nil {
		t.Fatalf("left trie invariant discrepancy: %v", d)
	}
	if d := sample.RightTrie.CheckInvariant(); d != nil {
		t.Fatalf("right trie invariant discrepancy: %v", d)
	}
	expected := New()
	for _, s := range setIntersect(sample.LeftTrie.List(), sample.RightTrie.List()) {
		expected.Add(s)
	}
	got := Intersect(sample.LeftTrie, sample.RightTrie)
	if d := got.CheckInvariant(); d != nil {
		t.Fatalf("got trie invariant discrepancy: %v", d)
	}
	if !Equal(expected, got) {
		t.Errorf("intersection of %v and %v: expected %v, got %v",
			sample.LeftTrie, sample.RightTrie, expected, got)
	}
}

type testIntersectTrie struct {
	LeftTrie  *Trie
	RightTrie *Trie
}

func testIntersectTrieFromJSON(srcJSON string) *testIntersectTrie {
	s := &testIntersectTrie{}
	if err := json.Unmarshal([]byte(srcJSON), s); err != nil {
		panic(err)
	}
	return s
}

var testIntersectJSONTries = []string{}
