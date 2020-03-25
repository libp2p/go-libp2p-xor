package key

import "bytes"

// Key is a vector of bits backed by a Go byte slice in big endian byte order and big-endian bit order.
type Key []byte

func (bs Key) BitAt(offset int) byte {
	if bs[offset/8]&(1<<(offset%8)) == 0 {
		return 0
	} else {
		return 1
	}
}

func (bs Key) BitLen() int {
	return 8 * len(bs)
}

func Equal(x, y Key) bool {
	return bytes.Equal(x, y)
}
