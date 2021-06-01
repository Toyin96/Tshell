package tShell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func execInput(input string) error{

	// remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// prepare the command to execute

}

func main(){
	for {
		// a reader where user input is being written on
		reader := bufio.NewReader(os.Stdin)

		/* reads the string until the first occurrence of a delimetre in the input
		which in this case is the newline symbol (\n). Nb: the delim is inclusive too in
		the input.
		 */
		input, err := reader.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
