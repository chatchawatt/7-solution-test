// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	var word string
	fmt.Scanln(&word)
	fmt.Println(decode(word))
}

func decode(encoded string) string {
	sequence := "0"
	prev := '0' // Initialize prev as an integer
	for _, symbol := range encoded {
		switch symbol {
		case 'L':
			prev--
			sequence += string(prev)
		case 'R':
			prev++
			sequence += string(prev)
		case '=':
			sequence += string(prev)
		}
	}
	for {
		if lessThanZero(sequence) {
			sequence = plusOne(sequence)
		} else {
			break
		}
	}
	return sequence
}
func plusOne(seq string) string {
	increased := ""
	for _, char := range seq {
		increased += string(char + 1)
	}
	return increased
}
func lessThanZero(seq string) bool {
	for _, char := range seq {
		if char < '0' {
			return true
		}
	}
	return false
}
