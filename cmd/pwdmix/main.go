package main

import (
	"fmt"
	"os"

	"pwdmix/internal/cli"
	"pwdmix/internal/pwd"
	"pwdmix/internal/utils"
)

func main() {
	opts, err := cli.ParseFlags()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	generatedPassword, err := pwd.PasswordGenerator(opts)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", *generatedPassword)
	utils.OverwriteMemory(generatedPassword)
}
