//James Kinsella - G00282261@gmit.ie
//Graph Theory Project 2018

package Graph_Project

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

func Poregtonfa(pofix string) *nfa {
	nfaStack := []*nfa{}
		case '.':
			
	}
	return ismatch
}