package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hampterl/Multitool/cli"
)

func Loop() {
	scanner := bufio.NewScanner(os.Stdin)

	//loop until exit
	for {
		//Start Menu Display
		Menu()

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			fmt.Print("|1|: Encode text to base64\n|2|: Decode base64 to text\n>")
			cli.UseBase64()
		case "2":
			fmt.Print("|1|: Encode text to hex\n|2|: Decode hex to text\n>")
			cli.UseHex()
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
		scanner.Scan()
	}
}
