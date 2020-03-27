package key

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Key is a vector of bits backed by a Go byte slice in big endian byte order and big-endian bit order.
type Key []byte

func (k Key) BitAt(offset int) byte {
	if k[offset/8]&(byte(1)<<(offset%8)) == 0 {
		return 0
	} else {
		return 1
	}
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
