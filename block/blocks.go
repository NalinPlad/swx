package block

import (
    "encoding/hex"
    "bytes"
)

var(
    Reverse Block = Block {
        Name: "Reverse",
        Description: "Reverses String",
        Out: 
        func(i string) string {
            runes := []rune(i)
            for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
                runes[i], runes[j] = runes[j], runes[i]
            }
            return string(runes)
        },
    }

    ToHex Block = Block {
        Name: "To Hex",
        Description: "Converts String to Hex",
        Out:
        func(i string) string {
            var hx = hex.EncodeToString([]byte(i))
            var buffer bytes.Buffer
            var l_1 = len(hx) - 1
            for i,rune := range hx {
               buffer.WriteRune(rune)
               if i % 2 == 1 && i != l_1  {
                  buffer.WriteRune(' ')
               }
            }
            return buffer.String()
        },
    }

)
