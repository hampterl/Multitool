package encodeDecode

import "encoding/base64"

func Base64Encode(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	return encoded
}

func Base64Decode(input string) string {
	decoded, _ := base64.StdEncoding.DecodeString(input)

	return string(decoded)
}
