package token

type Type int

// Token types
const (
	Text Type = iota
	EOF
	EOL
	Star
	Tick
	Colon
	Space
	Dot
	OBrack
	CBrack
	OABrack
	CABrack
	Tilda
	Score
	Plus
	Num
)
const (
	numType     = "NUMBER_TOKEN"
	plusType    = "PLUS_TOKEN"
	scoreType   = "SCORE_TOKEN"
	tildaType   = "TILDA_TOKEN"
	oabrackType = "OPEN_ANGLE_BRACKET_TOKEN"
	cabrackType = "CLOSED_ANGLE_BRACKET_TOKEN"
	obrackType  = "OPEN_BRACKET_TOKEN"
	cbrackType  = "CLOSED_BRACKET_TOKEN"
	dotType     = "DOT_TOKEN"
	spaceType   = "SPACE_TOKEN"
	colonType   = "COLON_TOKEN"
	tickType    = "TICK_TOKEN"
	starType    = "STAR_TOKEN"
	eolType     = "EOL_TOKEN"
	eofType     = "EOF_TOKEN"
	textType    = "TEXT_TOKEN"
)
