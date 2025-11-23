package main

import "github.com/01-edu/z01"

func main() {
	for ik := 'a'; ik <= 'z'; ik++ {
		z01.PrintRune(ik)
	}
	z01.PrintRune('\n')
}
