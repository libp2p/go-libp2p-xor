package key

import (
	"crypto/sha256"
	"math/big"
	"testing"

	kbucket "github.com/libp2p/go-libp2p-kbucket"
	ks "github.com/libp2p/go-libp2p-kbucket/keyspace"
)

func TestKbucketConversion(t *testing.T) {
	kbA, kbB := kbucket.ConvertKey("a"), kbucket.ConvertKey("b")
	ksA, ksB := ks.Key{Space: ks.XORKeySpace, Bytes: kbA}, ks.Key{Space: ks.XORKeySpace, Bytes: kbB}
	keyA, keyB := KbucketIDToKey(kbA), KbucketIDToKey(kbB)
	if DistInt(keyA, keyB).Cmp(ks.XORKeySpace.Distance(ksA, ksB)) != 0 {
		t.Errorf("key distance differs from kbucket distance")
	}
}

func TestBitEndiannes(t *testing.T) {
	r0 := sha256.Sum256([]byte("a")) // random bytes
	r := r0[:]
	key := KbucketIDToKey(r) // random key
	z := big.NewInt(0)
	for i := 0; i < key.BitLen(); i++ {
		// big.Int -> first byte most significant
		// big.Int -> first bit least significant
		z.SetBit(z, i, uint(key.BitAt(key.BitLen()-1-i)))
	}
	if key.NormInt().Cmp(z) != 0 {
		t.Errorf("bit-reconstructed norms differ: %v and %v", key.NormInt(), z)
	}
}

func TestKeyString(t *testing.T) {
	key := Key{0x05, 0xf0}
	if key.BitString() != "1111000000000101" {
		t.Errorf("unexpected bit string: %s", key.BitString())
	}
}

func TestBitAt(t *testing.T) {
	key := Key{0x21, 0x84}
	switch {
	case key.BitAt(0) != 1:
		t.Errorf("bit 0 flipped")
	case key.BitAt(4+1) != 1:
		t.Errorf("bit 5 flipped")
	case key.BitAt(8) != 0:
		t.Errorf("bit 8 flipped")
	case key.BitAt(8+2) != 1:
		t.Errorf("bit 10 flipped")
	case key.BitAt(8+4+3) != 1:
		t.Errorf("bit 15 flipped")
	}
}
