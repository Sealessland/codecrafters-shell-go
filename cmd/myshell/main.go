package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Trim the newline character
		command = strings.TrimSpace(command)

		switch {
		case command == "exit 0":
			os.Exit(0)
		case strings.HasPrefix(command, "echo "):
			// Print the rest of the command after "echo "
			fmt.Println(command[5:])
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
