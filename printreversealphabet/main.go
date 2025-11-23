package main

import "github.com/01-edu/z01"

func main() {
	for oga := 'z'; oga >= 'a'; oga-- {
		z01.PrintRune(oga)
	}
	z01.PrintRune('\n')
}
