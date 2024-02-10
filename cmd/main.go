package main

import (
	"fmt"
	"os"
)

func main(){

	userInput := os.Args[1:]

	if(len(userInput) > 0){
		fmt.Println("Hello, ", userInput[0]);
		return;
	}

	fmt.Println("Hello, World!");


}