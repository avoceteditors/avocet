package token

import "fmt"

func (t Type) String() string {
	switch t {
	case Text:
		return textType
	case EOF:
		return eofType
	case EOL:
		return eolType
	case Star:
		return starType
	case Tick:
		return tickType
	case Colon:
		return colonType
	case Space:
		return spaceType
	case Dot:
		return dotType
	case OBrack:
		return obrackType
	case CBrack:
		return cbrackType
	case OABrack:
		return oabrackType
	case CABrack:
		return cabrackType
	case Tilda:
		return tildaType
	case Score:
		return scoreType
	case Plus:
		return plusType
	}

	return fmt.Sprintf("UNKNOWN_TYPE(%d)", t)
}
