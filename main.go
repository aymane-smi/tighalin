package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/aymane-smi/tighalin/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Tighalin language REPL>>>>")
	repl.Start(os.Stdin, os.Stdout, user.Username)
}
