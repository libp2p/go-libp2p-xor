package key

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"math/bits"
	"strings"

	kbucket "github.com/libp2p/go-libp2p-kbucket"
)

// TODO: Find a way to prohibit casting bytes to Key.
func KbucketIDToKey(id kbucket.ID) Key {
	return Key(reverseBytesBits(id))
}

// Key is a vector of bits backed by a Go byte slice in big endian byte order and big-endian bit order.
// First byte is most significant.
// First bit in each byte is most significant.
type Key []byte

// reverseBytesBits reverses the bit-endianness of each byte in a slice.
func reverseBytesBits(blob []byte) []byte {
	r := make([]byte, len(blob))
	for i := range blob {
		r[i] = bits.Reverse8(blob[i])
	}
	return r
}

func (k Key) BitAt(offset int) byte {
	if k[offset/8]&(byte(1)<<(offset%8)) == 0 {
		return 0
	} else {
		return 1
	}
}

func (k Key) NormInt() *big.Int {
	return big.NewInt(0).SetBytes(reverseBytesBits(k))
}

func (k Key) BitLen() int {
	return 8 * len(k)
}

func (k Key) String() string {
	b, _ := json.Marshal(k)
	return string(b)
}

func (k Key) BitString() string {
	s := make([]string, len(k))
	for i, b := range k {
		s[len(k)-i-1] = fmt.Sprintf("%08b", b)
	}
	return strings.Join(s, "")
}

func Equal(x, y Key) bool {
	return bytes.Equal(x, y)
}

func Xor(x, y Key) Key {
	z := make([]byte, len(x))
	for i := range x {
		z[i] = x[i] ^ y[i]
	}
	return z
}

func DistInt(x, y Key) *big.Int {
	return Xor(x, y).NormInt()
}
