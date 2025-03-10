package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findExecutable(cmd string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, ":")
	for _, dir := range paths {
		fullPath := filepath.Join(dir, cmd)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, true
		}
	}
	return "", false
}
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
		case strings.HasPrefix(command, "type "):
			switch command[5:] {
			case "echo":
				fmt.Println("echo is a shell builtin")
			case "type":
				fmt.Println("type is a shell builtin")
			case "exit":
				fmt.Println("exit is a shell builtin")
			default:
				if _, found := findExecutable(command[5:]); found {
					fmt.Println(command[5:], "is", command[5:])
				} else {
					fmt.Println(command[5:], ": not found")
				}
				fmt.Println(command[5:] + ": not found")
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
