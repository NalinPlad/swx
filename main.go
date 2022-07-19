package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"

    "swx/block"
)


const stack_keyword = "<!>"


func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println(`
(v02)───────────────┬─►
   │ Welcome to SWX │
 ◄─┴────────────────(golang)
 `)
    
    // [1]   normal
    // [-1]  stack
    var state int = 1

    var stack [100]block.Block
    stack[0] = blocks.Reverse

    for {
        if state == -1{
            fmt.Print("stack editor> ")
        } else {
            fmt.Print("swx> ")
        }
        inp, _ := reader.ReadString('\n')
        inp = strings.Replace(inp, "\n", "", -1)

        // commands
        if inp == stack_keyword{
            // toggle state
            state = 0-state
        }

        if state == -1{
            // stack editor code
        } else{
            // normal mode code
            // var out string = inp            
            fmt.Println(inp)
        }
        
    }
}
