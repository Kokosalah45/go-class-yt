package main

import (
	"bufio"
	"fmt"
	compositetypes "go-class/composite-types"
	"go-class/file"
	"go-class/hello"
	"os"
)

func main(){
	
	option := os.Args[1]
	userInput := os.Args[2:]

	


	if option == "--hello"{
		fmt.Print(hello.Say(userInput)) ;		
		return
	}

	if option == "--replace" {
		old , new := os.Args[2] , os.Args[3]

		scanner := bufio.NewScanner(os.Stdin)
		
		fmt.Print(file.Replace(old, new, scanner))

		return
	}

	if option == "--composite-types" {
		compositetypes.Run()
	}

	if option == "--occurence" {
		
		
		scanner := bufio.NewScanner(os.Stdin)

		scanner.Split(bufio.ScanWords)

		fmt.Print(file.Occurence(scanner))

		return
	}





	
}