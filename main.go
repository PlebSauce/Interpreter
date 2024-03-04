package main

import (
	"Projects/Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

// when implementing in a different language like C, we need to make our own garbage collector
// Go has a GC built in for us, research this later
func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
