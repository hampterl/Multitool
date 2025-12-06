package repl

import (
	"fmt"

	"github.com/fatih/color"
)

func Menu() {
	//===============================================================================
	//Colors are a mess because I am still trying to figure out wich color looks best
	//===============================================================================

	blue := color.New(color.FgBlue).SprintFunc() //blue
	red2 := color.New(color.FgRed).SprintFunc()  //red
	//blue := color.New(color.FgGreen).SprintFunc() //green
	//blue := color.New(color.FgWhite).SprintFunc()  //white
	red := color.New(color.FgMagenta).SprintFunc() // puple
	//blue := color.New(color.FgCyan).SprintFunc()   //hellblau

	fmt.Println(
		"\n\n",
		blue("                   _____              ____			   --------------------------\n"),
		blue("                 /       \\          /      \\ 				\n"),
		blue("               /            \\ ____ /         ___ \\   /			")+red("|1|:")+" Encode/Decode\n",
		blue("             / ___      ____ ___      ____  ==   _\\_/_			")+red("|2|:")+" Encrypt/Decrypt\n",
		blue("          __  |___||   | |  |___||   | |            | |		")+red("|3|:")+" Coming Soon\n",
		blue("        / __  |    |___| |  |    |___| |    ==   _____			")+red("|4|:")+" Coming Soon\n",
		blue("       | |  \\-----------------------------   ___  / \\ 			")+red("|5|:")+" Coming Soon\n",
		blue("       |_|    \\                ____         /    /   \\ 		")+red("|0|:")+" Exit\n",
		blue("                \\__        __/      \\ ____ /				\n"),
		blue("                   _______					   --------------------------\n"),
		red2("\n	  	by hampterl   |   Version 1.0\n"),
		"\n",
		"-------------------------------------------------------------------------------------------\n")
	fmt.Print("> ")
}
