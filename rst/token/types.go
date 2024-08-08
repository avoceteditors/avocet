package token

import "fmt"

type Type int

const (
	Unset Type = iota
	EOL
	EOF
	Empty

	ParaStart
	ParaEnd

	UList
	OList
	CList
	Indent
)

const (
	unsetType = "UNSET_TOKEN"
	eolType   = "ENDOFLINE_TOKEN"
	eofType   = "ENDOFFILE_TOKEN"
	emptyType = "EMPTY_TOKEN"

	indentType = "INDENT_TOKEN"

	paraStartType = "PARAGRAPH_START_TOKEN"
	paraEndType   = "PARAGRAPH_END_TOKEN"

	ulistType = "ITEMIZED_LIST_TOKEN"
	olistType = "ENUMERATED_LIST_TOKEN"
	clistType = "CHECKLIST_TOKEN"
)

func (t Type) String() string {
	switch t {
	case Indent:
		return indentType
	case UList:
		return ulistType
	case OList:
		return olistType
	case CList:
		return clistType
	case Empty:
		return emptyType
	case ParaEnd:
		return paraEndType
	case ParaStart:
		return paraStartType
	case EOF:
		return eofType
	case EOL:
		return eolType
	case Unset:
		return unsetType
	default:
		return fmt.Sprintf("UNKNOWN_TOKEN_TYPE_%d", t)
	}
}
