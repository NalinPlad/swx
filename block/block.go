package block

import (
    "strings"
)

type Block struct {
	Name        string
	Description string
	Out         func(string) string
}

// Generate ID from string (parse input)
func Get_id(s string) string {
	
	id_split := strings.Split(s, "'")
	
	if len(id_split) < 2 {
		return "###invsyn"
	}

	id := strings.Replace(strings.ToLower(id_split[1]), " ", "_", -1)
    
    _, isBlock := Blocks[id]

    if isBlock {
		return id
	} else {
        return "###nf"
    }

}
