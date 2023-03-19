package main

import (
    "fmt"
    "strings"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    s = strings.ReplaceAll(s, " ", "") // Remove spaces
    n := len(s)
    for i := 0; i < n/2; i++ {
        if s[i] != s[n-i-1] {
            return false
        }
    }
    return true
}

func main() {
    input := "A man a plan a canal Panama"
    if isPalindrome(input) {
        fmt.Printf("%q is a palindrome\n", input)
    } else {
        fmt.Printf("%q is not a palindrome\n", input)
    }
}
