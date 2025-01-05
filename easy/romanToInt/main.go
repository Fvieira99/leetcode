package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	total := 0
	smap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	cmap := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	pos := 0
	lastIndex := len(s) - 1

	for pos <= lastIndex {
		if pos < lastIndex {
			cstr := string(s[pos]) + string(s[pos+1])
			if value, found := cmap[cstr]; found {
				total += value
				pos += 2
				continue
			} else {
				total += smap[string(s[pos])]
				pos++
				continue
			}
		}
		total += smap[string(s[pos])]
		pos++
	}

	return total
}
