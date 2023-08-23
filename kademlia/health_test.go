package kademlia

import (
	crand "crypto/rand"
	"math/rand"
	"testing"

	"github.com/libp2p/go-libp2p-xor/key"
)

func TestTableHealthFromSets(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := randomTestTableHealthSubsetSamples(10, 100)
		report := TableHealthFromSets(s.Node, s.Contacts, s.Known)
		for _, b := range report.Bucket {
			if b.ActualUnknownContacts != 0 {
				t.Errorf("expecting no actual unknown contacts, got %d", b.ActualUnknownContacts)
			}
		}
	}
}

func randomTestTableHealthSubsetSamples(contactSize, knownSize int) *testTableHealthSample {
	known := make([]key.Key, knownSize)
	for i := range known {
		known[i] = randomKey(16)
	}
	contacts := make([]key.Key, contactSize)
	perm := rand.Perm(knownSize)
	for i := 0; i < contactSize; i++ {
		contacts[i] = known[perm[i]]
	}
	return &testTableHealthSample{
		Node:     randomKey(16),
		Contacts: contacts,
		Known:    known,
	}
}

func randomKey(size int) key.Key {
	k := make([]byte, size)
	crand.Read(k)
	return key.BytesKey(k)
}

type testTableHealthSample struct {
	Node     key.Key
	Contacts []key.Key
	Known    []key.Key
}
