package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    count := 0
    if f, err := os.Open("input8.txt"); err != nil {
        log.Fatal("can't open input7.txt")
    } else {
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            line := scanner.Text()
            escaped := escape(line)
            count += len(escaped) - len(line)
        }
    }
    fmt.Println(count)
}

func escape(line string) string {
    newstr := []byte{'"'}
    for _, ch := range line {
        if ch == '\\' || ch == '"' {
            newstr = append(newstr, '\\')
        }
        newstr = append(newstr, byte(ch))
    }
    newstr = append(newstr, '"')
    return string(newstr)
}
