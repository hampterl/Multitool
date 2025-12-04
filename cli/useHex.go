package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hampterl/Multitool/internal/createFiles"
	"github.com/hampterl/Multitool/internal/encodeDecode"
)

func UseHex() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Text() == "" {
		fmt.Println("No input provided")
	}

	enDecode := strings.TrimSpace(scanner.Text())

	for {
		switch enDecode {
		case "1":
			encodehex()
			return
		case "2":
			decodehex()
			return
		default:
			fmt.Println("Invalid option\n|1|: Encode text to hex\n|2|: Decode hex to text\n>")
			return
		}
	}
}

func encodehex() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Text to encode: ")
	scanner.Scan()

	fmt.Print("Encoded text: ")
	encodedText := encodeDecode.HexEncode(scanner.Text())
	fmt.Println(encodedText)

	fmt.Print("Save txt into file? (y/n): ")
	scanner.Scan()
	saveFile := strings.TrimSpace(scanner.Text())
	for {
		switch saveFile {
		case "y":
			fmt.Print("Filename: ")
			scanner.Scan()
			filename := strings.TrimSpace(scanner.Text())

			createFiles.Savetxt(filename, encodedText)
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

func decodehex() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Base64 text: ")
	scanner.Scan()

	fmt.Print("Decoded text: ")
	decodedText := encodeDecode.HexDecode(scanner.Text())
	fmt.Println(decodedText)

	fmt.Print("Save txt into file? (y/n): ")
	scanner.Scan()
	saveFile := strings.TrimSpace(scanner.Text())
	for {
		switch saveFile {
		case "y":
			fmt.Print("Filename: ")
			scanner.Scan()
			filename := strings.TrimSpace(scanner.Text())

			createFiles.Savetxt(filename, decodedText)
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
