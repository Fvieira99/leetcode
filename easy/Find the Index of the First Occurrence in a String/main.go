package main

func main() {
	println(strStr("mississippi", "issip"))
	println(strStr("mississippi", "issipi"))
	println(strStr("leetcode", "leeto"))
	println(strStr("a", "a"))
}

// MELHOR RESPOSTA Usando Slice idx:idx
// func strStr(haystack string, needle string) int {
//     for idx, _ := range haystack {
//         if idx + len(needle) > len(haystack) {
//             break
//         }
//         if haystack[idx:idx + len(needle)] == needle {
//             return idx
//         }
//     }

//     return -1
// }

func strStr(haystack string, needle string) int {
	hp := 0
	np := 0
	matchStartIndex := -1

	if len(needle) > len(haystack) {
		return matchStartIndex
	}

	for i := 0; i <= len(haystack)-1; i++ {
		if np > len(needle)-1 {
			return matchStartIndex
		}

		if haystack[i] != needle[0] {
			continue
		}
		hp, matchStartIndex = i, i
		for np <= len(needle)-1 && hp <= len(haystack)-1 {
			hc := haystack[hp]
			nc := needle[np]

			if hc != nc {
				np = 0
				matchStartIndex = -1
				break
			}
			np++
			hp++
		}
	}

	if np <= len(needle)-1 {
		return -1
	}

	return matchStartIndex
}
