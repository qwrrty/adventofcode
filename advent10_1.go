package main

import "fmt"

func LookAndSay(s string) string {
    var result []rune
    i := 0
    for i < len(s) {
        j := i
        for j < len(s) && s[j] == s[i] {
            j++
        }
        count := rune('0') + rune(j-i)
        digit := rune(s[i])
        result = append(result, count, digit)
        i = j
    }
    return string(result)
}

func main() {
    input := "1113122113"
    next := input
    for i := 0; i < 50; i++ {
        next = LookAndSay(next)
    }
    fmt.Println(len(next))
}
