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
	left.CheckInvariant()
	for _, r := range sample.RightKeys {
		right.Add(r)
	}
	right.CheckInvariant()
	for _, s := range setIntersect(sample.LeftKeys, sample.RightKeys) {
		expected.Add(s)
	}
	got := Intersect(left, right)
	got.CheckInvariant()
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

func testIntersectSampleFromJSON(srcJSON string) *testIntersectSample {
	s := &testIntersectSample{}
	if err := json.Unmarshal([]byte(srcJSON), s); err != nil {
		panic(err)
	}
	return s
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
