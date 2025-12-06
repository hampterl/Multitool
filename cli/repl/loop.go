package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hampterl/Multitool/cli/choose"
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
			fmt.Print("|1|: Base64\n|2|: Hex\n|0|: Back\n> ")
			choose.ChooseEnDecode()
		case "2":
			fmt.Print("|1|: Aes256\n|2|: Comming Soon\n|0|: Back\n> ")
			choose.ChooseEnDecryption()
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
		scanner.Scan()
	}
}
