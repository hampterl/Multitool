package encodeDecode

import "encoding/hex"

func HexEncode(input string) string {
	encoded := hex.EncodeToString([]byte(input))

	return encoded
}

func HexDecode(input string) string {
	decoded, _ := hex.DecodeString(input)

	return string(decoded)
}
