package main

import "strconv"

// Return the century of the input year. The input will always be a 4 digit string, so there is no need for validation.
//
// Examples
// "1999" --> "20th"
// "2011" --> "21st"
// "2154" --> "22nd"
// "2259" --> "23rd"
// "1124" --> "12th"
// "2000" --> "20th"
func whatCentury(year string) string {
    y, _ := strconv.Atoi(year)
    century := (y + 99) / 100

    suffix := "th"
    if century%100 < 11 || century%100 > 13 {
        switch century % 10 {
        case 1:
            suffix = "st"
        case 2:
            suffix = "nd"
        case 3:
            suffix = "rd"
        }
    }

    return strconv.Itoa(century) + suffix
}