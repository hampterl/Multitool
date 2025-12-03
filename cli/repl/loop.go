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
			fmt.Println("Text to encode: ")
			cli.UseBase64()
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
		scanner.Scan()
	}
}
