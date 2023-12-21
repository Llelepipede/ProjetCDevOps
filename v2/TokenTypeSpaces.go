package v2

type TokenTypeSpacesBehavior struct {
	Name   string
	Syntax []string
}

func (t *TokenTypeSpacesBehavior) Resolve(l *TLexer) {
	l.DEBUGLEXER("in resolve spaces")
	if (*l).Inquote() {
		findNameInEveryTokenType((*l).TriggerBy, Every).Resolve(l)

	} else {
		(*l).isSpaces = true
	}
	if t.Name == "\n" {
		(*l).line++
		(*l).position = 0
	}

}

func (t *TokenTypeSpacesBehavior) Get() []string {
	return append(t.Syntax, t.Name)
}

func (t *TokenTypeSpacesBehavior) InvolvedWith() []ITokenType {
	return []ITokenType{}
}

var (
	EMPTY = TokenTypeSpacesBehavior{
		Name: "",
		Syntax: []string{
			" ",
			"\t",
			"\r",
		},
	}
	RETURN = TokenTypeSpacesBehavior{
		Name: "\n",
		Syntax: []string{
			"\n",
		},
	}
	SELF = TokenTypeSpacesBehavior{
		Name:   "SELF",
		Syntax: []string{},
	}
)
