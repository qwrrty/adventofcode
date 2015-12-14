package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

// A Gate represents an operation that may be performed on up to two
// input wires: e.g. "a AND b", "px RSHIFT 7", "NOT bg".
// arg1 and arg2 may be the names of wires, integers (representing constant
// inputs) or the empty string to indicate no input on that wire.
//
// The "cache" field holds the computed value of this gate. It is initialized
// to -1, indicating that the gate's value has not been computed.
type Gate struct {
    arg1, op, arg2 string
    cache int
}

func newGate(arg1, op, arg2 string) *Gate {
    return &Gate{arg1, op, arg2, -1}
}

// A breadboard is the data structure representing the whole circuit.
// It maps each wire name to the gate that provides its input.
type Breadboard map[string]*Gate

func newBreadboard() Breadboard {
    b := make(Breadboard)
    return b
}

// Reset the state of the breadboard (wipe out all cached values)
func (b Breadboard) reset() {
    for _, gate := range b {
        gate.cache = -1
    }
}

// parseLine takes an input line and returns the gate representing
// the input circuit and the name of the wire which receives it.
// e.g. on the input line "rx AND g -> bf", the return values
// would be &Gate{"rx", "AND", "g", -1} and "bf".
func parseLine(s string) (gate *Gate, wire string) {
    f := strings.Fields(s)
    switch len(f) {
    case 5:
        gate = newGate(f[0], f[1], f[2])
    case 4:
        gate = newGate(f[1], f[0], "")
    case 3:
        gate = newGate(f[0], "", "")
    }
    wire = f[len(f)-1]
    return
}

// evaluate recursively evaluates an input within the context of
// a breadboard circuit.  The input may be an empty string (representing
// no input), a numeric constant, or the name of a wire in the breadboard.
//
func evaluate(wire string, b Breadboard) int {
    if wire == "" {
        return 0
    }
    // If "wire" is a numeric literal, return that value
    if n, err := strconv.Atoi(wire); err == nil {
        return n
    }
    // Compute the wire's value from the gate, if it has not
    // already been cached.
    if b[wire].cache < 0 {
        // We need to compute the wire's value from the gate.
        // If the gate operation is empty, then arg1 is the sole
        // input and is a single wire or voltage.
        p := evaluate(b[wire].arg1, b)
        q := evaluate(b[wire].arg2, b)
        switch b[wire].op {
        case "":
            b[wire].cache = p
        case "NOT":
            b[wire].cache = ^p
        case "AND":
            b[wire].cache = p & q
        case "OR":
            b[wire].cache = p | q
        case "LSHIFT":
            b[wire].cache = p << uint16(q)
        case "RSHIFT":
            b[wire].cache = p >> uint16(q)
        }
    }
    return b[wire].cache
}

func main() {
    breadboard := newBreadboard()
    
    // Read the contents of input7.txt. Parse each line
    // into a gate and a target wire, and add that connection
    // to the breadboard.
    if f, err := os.Open("input7.txt"); err != nil {
        log.Fatal("can't open input7.txt")
    } else {
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
            gate, wire := parseLine(scanner.Text())
            breadboard[wire] = gate
        }
    }
    
    wire_a := evaluate("a", breadboard)
    fmt.Printf("a = %d\n", wire_a)
    
    // Now hardwire b to this signal, reset the board, and recalculate.
    breadboard["b"] = newGate(strconv.Itoa(wire_a), "", "")
    breadboard.reset()
    
    new_wire_a := evaluate("a", breadboard)
    fmt.Printf("a = %d\n", new_wire_a)
}
