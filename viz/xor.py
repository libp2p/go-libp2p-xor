import base64


def bits_in_byte(byte):
    return [
        byte & 0x1 != 0,
        byte & 0x2 != 0,
        byte & 0x4 != 0,
        byte & 0x8 != 0,
        byte & 0x10 != 0,
        byte & 0x20 != 0,
        byte & 0x40 != 0,
        byte & 0x80 != 0,
    ]


class Key(bytes):
    def to_float(self):
        f = 0.0
        s = 1.0
        for byte in self:
            for bit in bits_in_byte(byte):
                s /= 2.0
                if bit:
                    f += s
        return f


def xor_key(x: Key, y: Key):
    return Key(bytes([x[k] ^ y[k] for k in range(len(x))]))


def key_from_base64_optional(s: str):
    return key_from_base64(s) if s else None


def key_from_base64(s: str):
    return Key(base64.b64decode(s))
