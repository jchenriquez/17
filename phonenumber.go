package phonenumber

import (
	"fmt"
)

func letterCombinationsHelper(digits string, digitsMap map[byte][]byte, index int, combinations []string) []string {
	if index >= len(digits) {
		return combinations
	}

	nCombinations := make([]string, 0, len(combinations)*3)
	digit := digits[index] - '0'

	possibleCombs, in := digitsMap[digit]

	if !in {
		return []string{}
	}

	for _, letter := range possibleCombs {
		if len(combinations) == 0 {
			nCombinations = append(nCombinations, fmt.Sprintf("%c", letter))
		} else {
			for _, comb := range combinations {
				nCombinations = append(nCombinations, fmt.Sprintf("%s%c", comb, letter))
			}
		}
	}
	return letterCombinationsHelper(digits, digitsMap, index+1, nCombinations)
}

// LetterCombinations will return all letter combinations that can be produced by the provided digits
func LetterCombinations(digits string) []string {
	lettersMap := map[byte][]byte{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}
	for i := 0; i < len(digits); i++ {
		inDigit := digits[i] - '0'
		if inDigit < 2 || inDigit > 9 {
			return []string{}
		}
	}
	return letterCombinationsHelper(digits, lettersMap, 0, []string{})
}
