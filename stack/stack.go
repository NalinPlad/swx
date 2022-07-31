package stack

import (
	"swx/block"
	"fmt"
)

type Stack struct {
    Blocks	[]block.Block
}

func ModStack(s Stack, t string, inp string) Stack {
	
	switch t {
	case "add":
		id := block.Get_id(inp)
		// doesent exist
		if id == "###nf" {
			fmt.Printf("%s <- Block not found\n", inp)
		// invalid syntax
		} else if id == "###invsyn" {
			fmt.Println("Invalid syntax. EG: > add 'reverse'")
		// add to stack
		} else {
			_s := Stack{
				Blocks: append(s.Blocks, block.Blocks[id]),
			}
			return _s
		}
	
	}
	
	// Empty return because of quirk in compiler
	// https://github.com/golang/go/issues/39268
	return Stack{
		Blocks:	[]block.Block{},
	}

}
