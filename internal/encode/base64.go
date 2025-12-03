package encode

import "encoding/base64"

func Base64Encode(input string) string {
	encrypted := base64.StdEncoding.EncodeToString([]byte(input))

	return encrypted
}
