package main

import (
	"fmt"
	hello "go-class"
	"os"
)

func main(){
	
	userInput := os.Args[1:]
	fmt.Print(hello.Say(userInput)) ;
	return;
	
}