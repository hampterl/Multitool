package createFiles

import "os"

func Savetxt(filename string, content string) {
	os.WriteFile(filename, []byte(content), 0644)
}
