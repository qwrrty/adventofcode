package main

// --- Day 5: Doesn't He Have Intern-Elves For This? ---
//
// Santa needs help figuring out which strings in his text file are
// naughty or nice.
//
// A nice string is one with all of the following properties:
//
//   - It contains at least three vowels (aeiou only), like aei, xazegov,
//     or aeiouaeiouaeiou.
//   - It contains at least one letter that appears twice in a row, like
//     xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
//   - It does not contain the strings ab, cd, pq, or xy, even if they
//     are part of one of the other requirements.
//
// For example:
//
//   - ugknbfddgicrmopn is nice because it has at least three vowels
//     (u...i...o...), a double letter (...dd...), and none of the
//     disallowed substrings.
//   - aaa is nice because it has at least three vowels and a double
//     letter, even though the letters used by different rules overlap.
//   - jchzalrnumimnmhp is naughty because it has no double letter.
//   - haegwjzuvuyypxyu is naughty because it contains the string xy.
//   - dvszwmarrgswjxmb is naughty because it contains only one vowel.
//
// How many strings are nice?
//
// --- Part Two ---
//
// Realizing the error of his ways, Santa has switched to a better model
// of determining whether a string is naughty or nice. None of the old
// rules apply, as they are all clearly ridiculous.
//
// Now, a nice string is one with all of the following properties:
//
//   - It contains a pair of any two letters that appears at least twice
//     in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa),
//     but not like aaa (aa, but it overlaps).
//   - It contains at least one letter which repeats with exactly one
//     letter between them, like xyx, abcdefeghi (efe), or even aaa.
//
// For example:
//
//   - qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice
//     (qj) and a letter that repeats with exactly one letter between
//     them (zxz).
//   - xxyxx is nice because it has a pair that appears twice and a letter
//     that repeats with one between, even though the letters used by each
//     rule overlap.
//   - uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat
//     with a single letter between them.
//   - ieodomkazucvgmuy is naughty because it has a repeating letter with
//     one between (odo), but no pair that appears twice.
//
// How many strings are nice under these new rules?

import (
    "bufio"
    "log"
    "os"
    "fmt"
    "strings"
)

func main() {
    if f, err := os.Open("input5.txt"); err != nil {
        log.Fatal("could not open input5.txt")
    } else {
        var nice1, nice2 int
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            if isNice1(scanner.Text()) {
                nice1 += 1
            }
            if isNice2(scanner.Text()) {
                nice2 += 1
            }
        }
        fmt.Println(nice1)
        fmt.Println(nice2)
    }
}

func hasThreeVowels(s string) bool {
    vowels := 0
    for _, c := range s {
        if strings.ContainsRune("aeiou", c) {
            vowels += 1
            if vowels >= 3 {
                return true
            }
        }
    }
    return false
}

func hasDoubledLetter(s string) bool {
    for i := 0; i < len(s)-1; i++ {
        if s[i] == s[i+1] {
            return true
        }
    }
    return false
}

func hasEvilDigraph(s string) bool {
    return (strings.Contains(s, "ab") ||
    strings.Contains(s, "cd") ||
    strings.Contains(s, "pq") ||
    strings.Contains(s, "xy"))
}

func isNice1(s string) bool {
    return hasThreeVowels(s) && hasDoubledLetter(s) && !hasEvilDigraph(s)
}

func isNice2(s string) bool {
    hasDoubled := false
    hasSandwich := false
    for i := 0; i < len(s)-2 && !(hasDoubled && hasSandwich); i++ {
        digraph := s[i:i+2]
        if strings.Contains(s[i+2:], digraph) {
            hasDoubled = true
        }
        if s[i] == s[i+2] {
            hasSandwich = true
        }
    }
    return hasDoubled && hasSandwich
}
