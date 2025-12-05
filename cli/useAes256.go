package cli

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/hampterl/Multitool/internal/createFiles"
	"github.com/hampterl/Multitool/internal/encodeDecode"
)

func UseAes256() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Text() == "" {
		fmt.Println("No input provided")
	}

	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			encodeAes()
			return
		case "2":
			decodeAes()
			return
		default:
			fmt.Println("Invalid option\n|1|: Encode text to hex\n|2|: Decode hex to text\n>")
			return
		}
	}
}

func encodeAes() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Aes Key(256 byte): ")
	scanner.Scan()
	key := scanner.Text()

	fmt.Print("Text to encode: ")
	scanner.Scan()
	plaintext := strings.TrimSpace(scanner.Text())

	fmt.Print("Encoded text (hex): \n>")
	ciphertext, err := encodeDecode.EncryptAes(plaintext, []byte(key))
	if err != nil {
		fmt.Println(err)
	}

	hexEncodedAes := hex.EncodeToString(ciphertext)

	fmt.Println(hexEncodedAes)

	fmt.Print("Save txt into file? (y/n): ")
	scanner.Scan()
	saveFile := strings.TrimSpace(scanner.Text())
	for {
		switch saveFile {
		case "y":
			fmt.Print("Filename: ")
			scanner.Scan()
			filename := strings.TrimSpace(scanner.Text())

			createFiles.Savetxt(filename, hexEncodedAes)
			fmt.Println("File saved!\nPress enter to return to main menu...")
			return

		case "n":
			fmt.Println("NO file saved!\nPress enter to return to main menu...")
			return

		default:
			fmt.Print("Please enter y or n\nSave txt into file? (y/n): ")
			scanner.Scan()
		}
	}
}

func decodeAes() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Aes Key(32 byte): ")
	scanner.Scan()
	key := scanner.Text()

	fmt.Print("Encoded text (hex): ")
	scanner.Scan()
	hexciphertext := strings.TrimSpace(scanner.Text())

	ciphertext, _ := hex.DecodeString(hexciphertext)

	fmt.Print("Decoded text: ")
	plaintext, err := encodeDecode.DecryptAes(ciphertext, []byte(key))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(plaintext)

	fmt.Print("Save txt into file? (y/n): ")
	scanner.Scan()
	saveFile := strings.TrimSpace(scanner.Text())
	for {
		switch saveFile {
		case "y":
			fmt.Print("Filename: ")
			scanner.Scan()
			filename := strings.TrimSpace(scanner.Text())

			createFiles.Savetxt(filename, plaintext)
			fmt.Println("File saved!\nPress enter to return to main menu...")
			return

		case "n":
			fmt.Println("NO file saved!\nPress enter to return to main menu...")
			return

		default:
			fmt.Print("Please enter y or n\nSave txt into file? (y/n): ")
			scanner.Scan()
		}
	}
}
