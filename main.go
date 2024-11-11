package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		for _, env := range os.Environ() {
			fmt.Println(env)
		}
	}

	var i int
	for i = 0; i < len(args); i++ {
		switch args[i] {
		case "-u":
			if i+1 >= len(args) {
				fmt.Println("No environment variable to unset")
				os.Exit(1)
			}
			err := os.Unsetenv(args[i+1])
			if err != nil {
				fmt.Println("Error unsetting environment variable", args[i+1])
				os.Exit(1)
			}
			i++
			continue
		case "-i":
			os.Clearenv()
			continue
		}

		if !strings.Contains(args[i], "=") {
			break
		}

		varParts := strings.SplitN(args[i], "=", 2)
		if len(varParts) != 2 {
			fmt.Println("Invalid argument", args[i])
			os.Exit(1)
		}
		err := os.Setenv(varParts[0], varParts[1])
		if err != nil {
			fmt.Println("Error setting environment variable", varParts[0], varParts[1])
			os.Exit(1)
		}
	}

	if i >= len(args) {
		fmt.Println("No command to run")
		os.Exit(1)
	}

	cmd := exec.Command(args[i], args[i+1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command %s: %v\n", strings.Join(args[i:], " "), err)
		os.Exit(1)
	}
}
