package lexer

import "fmt"

// Actus reflects the lexer state with regard to how it processes lines of text.
type Actus int

// Actus state constants
const (
	OpenLine Actus = iota
	PassAct
	ParaFindEnd
)

// Actus state descriptive strings.
const (
	openLine    = "OPENLINE_ACTUS"
	passactLine = "PASSACT_ACTUS"
	paraFindEnd = "PARA_FIND_END_ACTUS"
)

// String returns the string representation of the Actus state.
func (t *Actus) String() string {
	switch *t {
	case ParaFindEnd:
		return paraFindEnd
	case OpenLine:
		return openLine
	case PassAct:
		return passactLine
	default:
		return fmt.Sprintf("UNKNOWN_ACTUS(%d)", *t)
	}
}
