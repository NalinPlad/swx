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
(v03)───────────────┬─►
   │ Welcome to SWX │
 ◄─┴────────────────(golang)
 `)
    
    // [1]   normal
    // [-1]  stack
    var state int = 1
    
    // Max stack size = 100 blocks
    var stack [100]block.Block

    // debug stack
    stack[0] = block.Reverse
    stack[1] = block.ToHex

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
            switch inp {
            case "list":
                for i, b := range stack {
                    if b.Name != ""{
                         fmt.Printf("[%d]✦ %s\n", i + 1, b.Name)
                    }
                }
            }
        } else{
            // normal mode code
            var out string = inp
            
            // i = index
            // b = block
            for i, b := range stack {
                _ = i
                if b.Name != ""{
                    out = b.Out(out)
                }
            }

            fmt.Println(out)
        }
        
    }
}
