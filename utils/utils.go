package utils

import (
    "strings"
)

// i: input
// s: search

func Parse(i string, s []string) bool {
	l := strings.ToLower(i)
   	o := false
	for _, x := range s {
       if strings.HasPrefix(l,x) {
            o = true
       }
    }

    return o
}
