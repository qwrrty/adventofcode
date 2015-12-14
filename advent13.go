package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

type Graph map[string]map[string]int

// Graph.Add adds the edge described by line to the graph.
// Example line: "AlphaCentauri to Snowdin = 66"
func (g Graph) Add(line string) {
    var (
        name1, gainlose, happiness, name2 string
    )
    rgx := regexp.MustCompile(
        `(\w+) would (gain|lose) (\d+).*sitting next to (\w+)\.`)
    if match := rgx.FindStringSubmatch(line); match == nil {
        log.Fatalf("could not parse %s", line)
    } else {
        name1 = match[1]
        gainlose = match[2]
        happiness = match[3]
        name2 = match[4]
    }
    dist, _ := strconv.Atoi(happiness)
    if gainlose == "lose" {
        dist = -dist
    }
    
    if g[name1] == nil {
        g[name1] = make(map[string]int)
    }
    g[name1][name2] = dist
}

func permute(list []string) (perms [][]string) {
    if len(list) <= 1 {
        perms = append(perms, list)
    } else {
        for i, item := range list {
            arg := make([]string, len(list)-1)
            copy(arg, list[0:i])
            copy(arg[i:], list[i+1:])
            subperms := permute(arg)
            for _, p := range subperms {
                p = append(p, item)
                perms = append(perms, p)
            }
        }
    }
    return perms
}

func calculateHappiness(graph Graph, seating []string) int {
    var happiness int
    for i, person := range seating {
        left  := seating[ (i-1 + len(seating)) % len(seating) ]
        right := seating[ (i+1)                % len(seating) ]
        happiness += graph[person][left] + graph[person][right]
    }
    return happiness
}

func main() {
    graph := make(Graph)

    if f, err := os.Open("input13.txt"); err != nil {
        log.Fatal("can't open input13.txt")
    } else {
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            graph.Add(scanner.Text())
        }
    }
    
    // Get a list of participants
    people := make([]string, 0)
    for p := range graph {
        people = append(people, p)
    }
    
    // Add yourself to the graph.
    for _, p := range people {
        graph.Add(
            fmt.Sprintf(
                "Tim would gain 0 happiness units by sitting next to %s.", p))
        graph.Add(
            fmt.Sprintf(
                "%s would gain 0 happiness units by sitting next to Tim.", p))
    }
    people = append(people, "Tim")

    perms := permute(people)
    best_seating := -1000000
    for _, p := range perms {
        // Calculate the seating value of this permutation
        seating_val := calculateHappiness(graph, p)
        if seating_val > best_seating {
            best_seating = seating_val
        }
    }
    fmt.Println(best_seating)
}
