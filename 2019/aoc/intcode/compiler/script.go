package compiler

const script = `
var a = input()
var loop = input()
for (var i = 0; i < loop; i++) {
	a = a * a
}
output(a)
halt()
`

func Compile() {

}

// type Declaration struct {
// 	Name  string `"var" @Ident`
// 	Value *Value `"=" @@`
// }

// type Value struct {
// 	Identifier    *string     `  @Ident`
// 	Number        *int        `| @Number`
// 	SubExpression *Expression `| @@`
// }

// type Assignment struct {
// 	Left  *string `@Ident`
// 	Right *Value  `"=" @@`
// }

// type Call struct {
// 	Name string `@Ident`
// 	Params []*Value `"(" (@@ [","])* ")"`
// }

// type Func struct {
// 	Name string `"func" @Ident`
// 	Params []*Value `"(" (@@ [","])* ")"`
// 	Expression *Expression
// }

// type Expression struct {
// 	Left *Value `@@`
// 	[]*OpAdd ``
// }
