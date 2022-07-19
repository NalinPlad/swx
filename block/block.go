package block

type Block struct {  
    Name            string
    Description     string
    Out             func(string) string
}
