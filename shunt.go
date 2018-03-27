//James Kinsella - G00282261@gmit.ie
//Graph Theory Project 2018

package main

import {
	"fmt"
}

func intopost(infix string) string {
	postfix := ""

	return postfix
}

func main() {
	// Answer ab.c*.
	fmt.PrintLn("Infix:	","a.b.c*")
	fmt.PrintLn("Postfix: ", intopost("a.b.c*"))

	// Answer abd|.*
	fmt.PrintLn("Infix:	","(a.(b|d))*")
	fmt.PrintLn("Postfix: ", intopost("(a.(b|d))*"))

	// Answer abd|.c*
	fmt.PrintLn("Infix:	","a.(b|d).c*")
	fmt.PrintLn("Postfix: ", intopost("a.(b|d).c*"))

	// Answer abb.+.c.
	fmt.PrintLn("Infix:	","a.(b.b)+.c*")
	fmt.PrintLn("Postfix: ", intopost("a.(b.b)+.c*"))



}
