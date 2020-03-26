package key

import "testing"

func TestKeyString(t *testing.T) {
	key := Key{0x05, 0xf0}
	println(key.String())
}
