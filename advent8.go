package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
    count := 0
    if f, err := os.Open("input8.txt"); err != nil {
        log.Fatal("can't open input7.txt")
    } else {
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            line := scanner.Text()
            parsed := unescape(line)
            count += len(line) - len(parsed)
        }
    }
    fmt.Println(count)
}

func unescape(line string) string {
    if line[0] != '"' {
        return ""
    }

    parsed := make([]byte, 0)
    for i := 1; line[i] != '"'; i++ {
        if line[i] == '\\' {
            switch line[i+1] {
            case '\\', '"':
                i++
                parsed = append(parsed, line[i])
            case 'x':
                if hexval, err := strconv.ParseInt(line[i+2:i+4], 16, 16); err == nil {
                    parsed = append(parsed, byte(hexval))
                    i += 3
                } else {
                    log.Fatalf("parsing hex escape %s: %s", line[i:i+4], err)
                }
            default:
                log.Fatalf("unrecognized escape %s", line[i:i+2])
            }
        } else {
            parsed = append(parsed, line[i])
        }
    }
    return string(parsed)
}
