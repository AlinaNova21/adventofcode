package assembler

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
)

// Comment is a comment
type Comment struct {
	Pos     lexer.Position
	Comment string `parser:"@Comment"`
}

// Label is a label, used to point to an address
type Label struct {
	Pos  lexer.Position
	Name string `parser:"@Ident ':'"`
}

// Mode is the ParameterMode
type Mode int

// Capture captures the Mode
func (m *Mode) Capture(v []string) error {
	*m = 0
	if v[0] == "%" {
		*m = 1
	}
	if v[0] == "@" {
		*m = 2
	}
	return nil
}

// Param is an Instruction parameter
type Param struct {
	Pos lexer.Position
	// Mode       *Mode
	Mode       Mode    `parser:"@( '@' | '%')?"`
	Number     *int    `parser:"( @Number"`
	Identifier *string `parser:"| @Ident )"`
	// Boolean    *Bool   `parser:"| @( "true" | "false" ) )"`
}

// Instruction is an IntCode instruction
type Instruction struct {
	Pos    lexer.Position
	Label  *Label   `parser:"@@?"`
	Data   string   `parser:"(  'DATA'"`
	Op     string   `parser:"|  @Ident )"`
	Params []*Param `parser:"  @@*"`
}

// Program is an IntCode program
type Program struct {
	Pos          lexer.Position
	Comments     []string       `parser:"( ( @Comment"`
	Instructions []*Instruction `parser:"  | ( @@ ))? EOL)*"`
}

// GetParser returns an IntCode parser for IntCode assembly
func GetParser(v interface{}) (*participle.Parser, error) {
	intcodeLexer := lexer.Must(ebnf.New(`
		Comment = ("#" | "//" ) { "\u0000"…"\uffff"-"\n"-"\r" } .
		Ident = (alpha | "_") { "_" | alpha | digit } .
		String = "\"" { "\u0000"…"\uffff"-"\""-"\\" | "\\" any } "\"" .
		Number = [ "-" | "+" ] ("." | digit) { "." | digit } .
		Punct = "!"…"/" | ":"…"@" | "["…` + "\"`\"" + ` | "{"…"~" .
		EOL = ( "\n" | "\r" ) { "\n" | "\r" }.
		Whitespace = ( " " | "\t" ) { " " | "\t" } .
		Op = alpha { alpha } .
		Alpha = alpha .
		AlphaNumeric = ( alpha | digit ) .
		alpha = "a"…"z" | "A"…"Z" .
		digit = "0"…"9" .
		any = "\u0000"…"\uffff" .
	`))
	return participle.Build(v,
		participle.Lexer(intcodeLexer),
		participle.UseLookahead(2),
		participle.CaseInsensitive("op"),
		participle.Elide("Whitespace"),
		participle.Upper("Op"),
	)
}
