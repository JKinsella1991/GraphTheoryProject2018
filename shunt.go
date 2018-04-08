//James Kinsella - G00282261@gmit.ie
//Graph Theory Project 2018

package Graph_Project



func Intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	pofix := []rune{}
	s := []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				pofix = append(pofix, s[len(s)-1])
				s = s[:len(s)-1]
			}
			s = s[:len(s)-1]
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {

				pofix = append(pofix, s[len(s)-1])
				s = s[:len(s)-1]
			}

			s = append(s, r)

		default:
			pofix = append(pofix, r)
		}
	}
	for len(s) > 0 {
		pofix = append(pofix, s[len(s)-1])
		s = s[:len(s)-1]
	}
	return string(pofix)
}

func StringTrim(s string) string {

	if len(s) > 0 {
		s = s[:len(s) - 2]
	}
	return s
}
