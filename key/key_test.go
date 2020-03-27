package key

import "testing"

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
