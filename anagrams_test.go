package pkg

import (
    "testing"
    "strings"
//    "fmt"
)

func AssertTrue(t *testing.T, condition bool, desc string) {
    if !condition {
        t.Error(desc)
    }
}

func anagramMatchingSeed(anagram, seed string) bool {
    if strings.Contains(strings.Join(anagrams(seed), "|"), anagram) {
        return true
    }
    
    return false
}

func TestOneChar(t *testing.T) {
    seed := "A"
    AssertTrue(t, anagramMatchingSeed("A", seed), "")
    
    seed = "B"
    AssertTrue(t, anagramMatchingSeed("B", seed), "")
}

func TestTwoChars (t *testing.T) {
    seed := "AB"
    AssertTrue(t, anagramMatchingSeed("AB", seed), "")
    AssertTrue(t, anagramMatchingSeed("BA", seed), "")
   // fmt.Print(anagrams("ABC"))
    seed = "BC"
    AssertTrue(t, anagramMatchingSeed("BC", seed), "")
    AssertTrue(t, anagramMatchingSeed("CB", seed), "")
    
    seed = "YZ"
    AssertTrue(t, anagramMatchingSeed("YZ", seed), "")
    AssertTrue(t, anagramMatchingSeed("ZY", seed), "")    
}

func TestThreeChars (t *testing.T) {
    seed := "ABC"
   // fmt.Print(anagrams("ABC"))
    AssertTrue(t, anagramMatchingSeed("ABC", seed), "")
    AssertTrue(t, anagramMatchingSeed("ACB", seed), "")
    AssertTrue(t, anagramMatchingSeed("BAC", seed), "")
    AssertTrue(t, anagramMatchingSeed("BCA", seed), "")

    seed = "ABCD"
    //fmt.Print(anagrams(seed))    
    AssertTrue(t, anagramMatchingSeed("ABCD", seed), "")
    AssertTrue(t, anagramMatchingSeed("ACBD", seed), "")
    AssertTrue(t, anagramMatchingSeed("BACD", seed), "")
    AssertTrue(t, anagramMatchingSeed("BCDA", seed), "")
      
}