package main

import "fmt"
import "strings"

func incrementPassword(s string) string {
	r := []rune(s)
	for i := len(r) - 1; i >= 0; i-- {
		r[i]++
		if r[i] <= 'z' {
			break
		}
		// roll over
		r[i] = 'a'
	}
	return string(r)
}

func hasRunOfThree(pw string) bool {
	// Check for runs of three
	for i := 0; i < len(pw)-2; i++ {
		if pw[i+1] == pw[i]+1 && pw[i+2] == pw[i]+2 {
			return true
		}
	}
    return false
}

func hasTwoDoubles(pw string) bool {
    var dblCount int
    for i := 0; i < len(pw)-1; i++ {
        if pw[i] == pw[i+1] {
            i++
            dblCount++
            if dblCount >= 2 {
                return true
            }
        }
    }
    return false
}

func hasIllegalLetter(pw string) bool {
    return strings.ContainsAny(pw, "iol")
}

func isLegalPassword(pw string) bool {
    return hasRunOfThree(pw) && hasTwoDoubles(pw) && !hasIllegalLetter(pw)
}

func nextLegalPassword(pw string) string {
    next := incrementPassword(pw)
    for !isLegalPassword(next) {
        next = incrementPassword(next)
    }
    return next
}

func main() {
	password := "vzbxkghb"
    next_pw := nextLegalPassword(password)
    next_next_pw := nextLegalPassword(next_pw)
    fmt.Println(next_pw)
    fmt.Println(next_next_pw)
}
