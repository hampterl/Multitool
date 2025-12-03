package createFiles

import "os"

func SaveBase64txt(filename string, content string) {
	os.WriteFile(filename, []byte(content), 0644)
}
