package block

func main(){
    Reverse := Block {
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

    _ = Reverse

}
