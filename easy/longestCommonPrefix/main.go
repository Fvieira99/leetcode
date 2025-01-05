package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	smallestWordLen := len(strs[0])
	var lcp string

	for _, currWord := range strs {
		currWordLen := len(currWord)
		if currWordLen < smallestWordLen {
			smallestWordLen = currWordLen
		}
	}

	for i := 0; i <= smallestWordLen-1; i++ {
		auxChar := ""
		for j, currWord := range strs {
			currWordChar := string(currWord[i])
			if j == 0 && auxChar == "" {
				auxChar = currWordChar
				continue
			}

			if i == 0 && currWordChar != auxChar {
				return ""
			}

			if i != 0 && currWordChar != auxChar {
				auxChar = ""
				return lcp
			}
		}
		lcp += auxChar
	}

	return lcp
}
