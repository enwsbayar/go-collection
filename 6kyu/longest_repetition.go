package main

type Result struct {
    C rune
    L int
}

func longestRepetition(s string) Result {
    if len(s) == 0 {
        return Result{}
    }

    best := Result{C: rune(s[0]), L: 1}
    currentChar := best.C
    currentLen := 1

    for _, r := range s[1:] {
        if r == currentChar {
            currentLen++
            if currentLen > best.L {
                best.L = currentLen
                best.C = currentChar
            }
        } else {
            currentChar = r
            currentLen = 1
        }
    }

    return best
}