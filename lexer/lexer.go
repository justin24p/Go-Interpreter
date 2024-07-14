package lexer


type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

// way in go to create an intialize a new instance of a struct 
// *Lexer points to struct above 
func New(input string) *Lexer {
	// the & means it will return a pointer to the newly struct instance of Lexer
	// not a copy therfore direct mainpulation and more efficient! 
	l := &Lexer{input: input}
	return l
}
// so readchar allows us to move through the input sourcecode
// example x + 5 lets say x gets indentified poistion would mark the start of x
/// read position would already be one ahead pointing to + 
// position is what will allow us to know were and until we should generate a token

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}