package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"swx/block"
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

	var stack = []block.Block{}

	for {
		if state == -1 {
			fmt.Print("stack editor✦> ")
		} else {
			fmt.Print("swx✦> ")
		}
		inp, _ := reader.ReadString('\n')
		inp = strings.Replace(inp, "\n", "", -1)

		// Toggle state
		if inp == stack_keyword {
			state = 0 - state
			continue
		} else if inp == "quit" || inp == "exit" {
			os.Exit(0)
		}

		if state == -1 {
			switch inp {

			case "list":
				for i, b := range stack {
					fmt.Printf("[%d]✦ %s\n", i+1, b.Name)
				}

			// If command is not one word, parse below
			default:
				if strings.HasPrefix(inp, "add ") {

					id := block.Get_id(inp)

					// doesent exist
					if id == "##nf" {
						fmt.Printf("Block -> %s <- not found\n", inp)
					// invalid syntax
					} else if id == "###ivnsyn" {
						fmt.Println("Invalid syntax. EG: > add 'reverse'")
					// add to stack
					} else {
						stack = append(stack, block.Blocks[id])
					}

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
			for _, b := range stack {
				out = b.Out(out)
			}

			fmt.Println(out)
		}

	}
}
