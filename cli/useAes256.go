package cli

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/hampterl/Multitool/internal/createFiles"
	"github.com/hampterl/Multitool/internal/crypto"
)

func UseAes256() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("|1|: Encrypt text with aes256\n|2|: Decrypt aes256 to text\n|3|: Encrypt File with aes256\n|0|: Return to main menu\n> ")

	scanner.Scan()

	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			encryptAes()
			return
		case "2":
			decryptAes()
			return
		case "3":
			encryptFile()
			return
		case "0":
			fmt.Println("Press enter to return to main menu...")
			return
		default:
			fmt.Print("Invalid option\n|1|: Encrypt text with aes256\n|2|: Decrypt aes256 to text\n|3|: Encrypt File with aes256\n|0|: Return to main menu\n> ")
			scanner.Scan()
			enDecode = strings.TrimSpace(scanner.Text())
			break
		}
	}
}

func encryptAes() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Aes Key(16, 24, 32 bytes): ")
	scanner.Scan()
	key := scanner.Text()

	fmt.Print("Text to encode: ")
	scanner.Scan()
	plaintext := strings.TrimSpace(scanner.Text())

	fmt.Print("Encoded text (hex): \n>")
	ciphertext, err := crypto.EncryptAes(plaintext, []byte(key))
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

func decryptAes() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Aes Key(16, 24, 32 bytes): ")
	scanner.Scan()
	key := scanner.Text()

	fmt.Print("Encoded text (hex): ")
	scanner.Scan()
	hexciphertext := strings.TrimSpace(scanner.Text())

	ciphertext, _ := hex.DecodeString(hexciphertext)

	fmt.Print("Decoded text: ")
	plaintext, err := crypto.DecryptAes(ciphertext, []byte(key))
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

func encryptFile() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Instructions in README.md!\n")
	fmt.Print("DISCLAIMER: This WILL encrypt the file you chose! IF you lose your key and overwrite your original," +
		" you will NOT be able to decrypt the file!\n" +
		"Never share your key with anyone! Make a backup before encrypting! Use at own Risk!\n")

	fmt.Print("Filepath: ")
	scanner.Scan()
	filepath := strings.TrimSpace(scanner.Text())

	fmt.Print("Save location: ")
	scanner.Scan()
	savepath := strings.TrimSpace(scanner.Text()) + ".enc"

	fmt.Print("Aes Key(16, 24, 32 bytes): ")
	scanner.Scan()
	key := scanner.Text()

	crypto.EncryptFile(filepath, savepath, key)

	fmt.Println("File encrypted! SAVE YOUR KEY!\nPress enter to return to main menu...")
}
