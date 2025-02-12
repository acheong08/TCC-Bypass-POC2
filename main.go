package main

import (
	"bufio"
	"fmt"
	mod "github.com/acheong08/GoShell/modules"
	"os"
)

func main() {
	var cmd string
	reader := bufio.NewReader(os.Stdin)
	for cmd != "exit\n" {
		fmt.Print("Enter command: ")
		cmd, _ = reader.ReadString('\n')
		if cmd == "screenshot\n" {
			mod.Screenshot()
			fmt.Println("Done")
		} else if cmd == "keyboard\n" {
			mod.Keyboard()
			fmt.Println("\nDone")
		} else {
			out, err := mod.Shell(cmd)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\nExec output: %s \n", out)
			}
		}
	}
}
