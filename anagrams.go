package pkg

import (
//    "fmt"
//    "strings"
)

func anagrams(seed string) []string {
    var result []string
    
    var f func(string, string) []string
    f = func (partialSeed string, prefix string) []string {
        for i := range partialSeed {
            l_prefix := prefix
		    if len(partialSeed) == 1 {
	            result = append(result, l_prefix + partialSeed)
	        } else {
	           l_prefix = l_prefix + string(partialSeed[i])
	           nextSeed := partialSeed[0:i] + partialSeed[i+1:] 
              	           
	           f(nextSeed, l_prefix)
	        }
	    }	    

	    return result
    }
    
    return f(seed, "")
}

