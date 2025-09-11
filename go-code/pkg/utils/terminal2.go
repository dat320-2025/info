package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunShell starts a simple interactive shell.
func RunShell() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("go-shell> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		if args[0] == "exit" {
			fmt.Println("Exiting shell.")
			break
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
