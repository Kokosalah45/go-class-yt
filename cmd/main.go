package main

import (
	"fmt"
	hello "go-class"
	"os"
)

func main(){

	userInput := os.Args[1:]

	if(len(userInput) > 0){
		fmt.Print(hello.Say(userInput)) ;
		return;
	}

	fmt.Println("Hello, World!");


}