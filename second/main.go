package main

import (
	"fmt"
	"os"
	"second/cmd"
)

func main() { 

		fmt.Println(cmd.Say(os.Args[1:]))
	
	
}