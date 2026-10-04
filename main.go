package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kakeetopius/mlang/repl"
)

func main() {
	user, err := user.Current()
	checkErr(err)

	fmt.Printf("Hello %s! This is mlang Programming language. Feel free to type in commands.\n\n", user.Name)

	repl.Start(os.Stdin, os.Stdout)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
