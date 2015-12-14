package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

type Edge struct {
    c1, c2 string
    distance int
}

const MAXDISTANCE = 1000000

type Graph map[string]map[string]int

// Graph.Add adds the edge described by line to the graph.
// Example line: "AlphaCentauri to Snowdin = 66"
func (g Graph) Add(line string) {
    var v1, v2 string
    var dist int
    
    rgx := regexp.MustCompile(`^(\S+) to (\S+) = (\d+)`)
    if match := rgx.FindStringSubmatch(line); match == nil {
        log.Fatalf("cannot parse \"%s\"", line)
    } else {
        v1, v2 = match[1], match[2]
        dist, _ = strconv.Atoi(match[3])
    }
    
    if g[v1] == nil {
        g[v1] = make(map[string]int)
    }
    g[v1][v2] = dist
    if g[v2] == nil {
        g[v2] = make(map[string]int)
    }
    g[v2][v1] = dist
}

func (g Graph) Distance(v1, v2 string) (int, bool) {
    if g[v1] != nil {
        d, ok := g[v1][v2]
        return d, ok
    }
    return 0, false
}

func (g Graph) ShortestPath(source string) int {
    unvisited := make(map[string]bool)
    for v := range g {
        unvisited[v] = true
    }
    delete(unvisited, source)
    
    current_node := source
    current_distance := 0

    // Keep finding the nearest unvisited neighbor and add it to the path
    // until no unvisited nodes are left.
    for len(unvisited) > 0 {
        best := ""
        best_leg := MAXDISTANCE
        for next := range unvisited {
            if next_leg, ok := g.Distance(current_node, next); ok {
                if next_leg < best_leg {
                    best = next
                    best_leg = next_leg
                }
            }
        }
        current_node = best
        current_distance += best_leg
        delete(unvisited, best)
    }
    return current_distance
}

func main() {
    graph := make(Graph)
    
    if f, err := os.Open("input9.txt"); err != nil {
        log.Fatal("can't open input9.txt")
    } else {
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            line := scanner.Text()
            graph.Add(line)
        }
    }

    sp := MAXDISTANCE
    for start := range graph {
        p := graph.ShortestPath(start)
        if p < sp {
            sp = p
        }
    }
    fmt.Println(sp)
}
