package utils

import (
	"os"
	"os/exec"
	"strings"
)

func ExecInput(input string) error{

	// remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// prepare the command to execute
	cmd := exec.Command(input)

	//set the correct out device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//execute the command and return the error
	return cmd.Run()

}
