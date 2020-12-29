package main

import "fmt"

func main() {
    // A string.
    value := "bird"

    // Part A: take substring with runes.
    // ... This handles any kind of rune in the string.
    runes := []rune(value)
    // ... Convert back into a string from rune slice.
    safeSubstring := string(runes[1:4])
    fmt.Println(" RUNE SUBSTRING:", safeSubstring)

    // Part B: take substring with direct byte slice.
    // ... This handles only ASCII correctly.
    asciiSubstring := value[1:3]
    fmt.Println("ASCII SUBSTRING:", asciiSubstring)
}
