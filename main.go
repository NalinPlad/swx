package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"swx/block"
	"swx/stack"
	"swx/utils"
)

const stack_keyword = "<!>"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(`
(v03)────✦─────────✦┬─►
  ✦│ Welcome✦to SWX │ ✦
 ◄─┴──✦──────────✦──(golang)
 `)

	// [1]   normal
	// [-1]  stack
	var state int = 1

	// Init empty stack
	var prog_stack = stack.Stack{
		Blocks: []block.Block{},
	}

	for {
		if state == -1 {
			fmt.Print("stack editor✦> ")
		} else {
			fmt.Print("swx✦> ")
		}
		// Get input
		inp, _ := reader.ReadString('\n')
		inp = strings.Replace(inp, "\n", "", -1)

		// Toggle state
		if inp == stack_keyword {
			state = 0 - state
			continue
		} else if inp == "quit" || inp == "exit" {
			os.Exit(0)
		}
		
		// Stack Editor mode
		if state == -1 {
			switch inp {

			case "list", "ls":
				for i, b := range prog_stack.Blocks {
					fmt.Printf("[%d]✦ %s\n", i+1, b.Name)
				}

			// If command is not one word, parse below
			default:

				// console
				if utils.Parse(inp, []string{"add", "a", "+"}) {
					prog_stack = stack.ModStack(prog_stack, "add", inp)
				}

			} // end switch
		} else {
			if inp == "" {
				continue
			}

			// normal mode code
			var out string = inp

			// i = index
			// b = block
			for _, b := range prog_stack.Blocks {
				out = b.Out(out)
			}

			fmt.Println(out)
		}

	}
}
