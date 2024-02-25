package main

import (
	"bufio"
	"fmt"
	"go-class/hello"
	"go-class/replace"
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

		fmt.Print(replace.Replace(old, new, scanner))

		return
	}





	
}