package main

import (
	"fmt"
	"github.com/962n/writing-an-interpreter-in-go/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!!\n", u.Name)
	repl.Start(os.Stdin, os.Stdout)
}
