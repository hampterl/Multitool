package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hampterl/Multitool/internal/createFiles"
	"github.com/hampterl/Multitool/internal/encode"
)

func UseBase64() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Text() == "" {
		fmt.Println("No input provided")
	}

	fmt.Print("Encoded text: ")
	fmt.Println(encode.Base64Encode(scanner.Text()))

	fmt.Println("Save txt into file? (y/n): ")
	scanner.Scan()
	saveFile := strings.TrimSpace(scanner.Text())

	switch saveFile {
	case "y":
		fmt.Print("Filename: ")
		scanner.Scan()
		filename := strings.TrimSpace(scanner.Text())

		createFiles.SaveBase64txt(filename, encode.Base64Encode(scanner.Text()))
		fmt.Println("File saved!\nPress enter to return to main menu...")
		return

	case "n":
		fmt.Println("NO file saved!\nPress enter to return to main menu...")
		return

	default:
		fmt.Println("Please enter y or n")
		scanner.Scan()
	}
}
