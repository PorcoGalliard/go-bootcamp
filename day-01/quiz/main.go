package main

import (
	"fmt"
	"strings"
	"unicode"
)
func main() {
	findDivisor(6)
	findDivisor(24)
	findDivisor(7)

	extractDigit(12234)
	extractDigit(5432)
	extractDigit(1278)

	segitigaV1(5)

	segitigaV2(5)

	piramida(8)

	awalAkhirIncrement(9)

	ganjilGenapDash(9)

	fmt.Println(isPalindrome("Kasur ini rusak"))
	fmt.Println(isPalindrome("tamaT"))           
	fmt.Println(isPalindrome("Aku Usa"))
	
	fmt.Println(reverseString("ABCD"))  
	fmt.Println(reverseString("tamaT")) 
	fmt.Println(reverseString("XYnb")) 

	fmt.Println(checkBraces("()[]{}"))
	fmt.Println(checkBraces("(]"))
	fmt.Println(checkBraces("({[]})"))    
	fmt.Println(checkBraces("((("))
	
	fmt.Println(isPalindromeNumber(121))  
	fmt.Println(isPalindromeNumber(1221)) 
	fmt.Println(isPalindromeNumber(123))  
	fmt.Println(isPalindromeNumber(10))   
}

func findDivisor(n int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func extractDigit(n int) {
	for n > 0 {
		fmt.Print(n%10, " ")
		n /= 10
	}
	fmt.Println()
}

func segitigaV1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func segitigaV2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func piramida(n int) {
	for i := n; i >= 1; i-- {

		for j := i; j >= 1; j-- {
			fmt.Print(j, " ")
		}

		for j := 2; j <= i; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}

func awalAkhirIncrement(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%2 != 0 {
				fmt.Print(i, " ")
			} else {
				fmt.Print(n-i+1, " ")
			}
		}
		fmt.Println()
	}
}

func ganjilGenapDash(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (i%2 != 0 && j%2 != 0) || (i%2 == 0 && j%2 == 0) {
				fmt.Print("- ")
			} else {
				fmt.Print(j, " ")
			}
		}
		fmt.Println()
	}
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	var cleaned []rune
	for _, ch := range s {
		if !unicode.IsSpace(ch) {
			cleaned = append(cleaned, ch)
		}
	}

	i, j := 0, len(cleaned)-1
	for i < j {
		if cleaned[i] != cleaned[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func reverseString(s string) string {
	r := []rune(s)
	i, j := 0, len(r)-1

	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return string(r)
}

func checkBraces(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if top != pairs[ch] {
				return false
			}
		}
	}

	return len(stack) == 0
}

func isPalindromeNumber(n int) bool {
	if n < 0 {
		return false
	}

	original := n
	reversed := 0

	for n > 0 {
		reversed = reversed*10 + (n % 10)
		n /= 10
	}

	return original == reversed
}




