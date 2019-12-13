package compiler

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"
)

const testScript = `package main
func main() {
	abc := input()
	output(abc)
}

func test(a, b int, c *int) int {
}
`

func TestCompile(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test", testScript, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	f, _ := os.Create("/tmp/ast")
	defer f.Close()
	out := make([]string, 0)
	ast.Fprint(f, fset, file, nil)
	callMap := map[string][]string{}
	callMap["input"] = []string{"IN %0"}
	callMap["output"] = []string{"OUT %0"}
	initial := []string{
		"REL 1000",
		"JZ %0 __func_main",
	}
	out = append(out, initial...)
	ast.Inspect(file, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.FuncDecl:
			_ = v
		case *ast.CallExpr:
			fname := v.Fun.(*ast.Ident).Name
			if fn, ok := callMap[fname]; ok {
				out = append(out, fn...)
				return false
			}
		case *ast.Ident:
			kind := "ident"
			if v.Obj != nil {
				kind = v.Obj.Kind.String()
			}
			name := fmt.Sprintf("__%s_%s:", kind, v.Name)
			out = append(out, name)
		}
		return true
	})
	fmt.Println(strings.Join(out, "\n"))
}

func call(name string, args ...string) []string {
	codeSize := len(args) + 4
	stackSize := len(args) + 1
	out := make([]string, codeSize)
	out[0] = fmt.Sprintf("REL -%d", stackSize)
	for i, a := range args {
		out[i+1] = fmt.Sprintf("ADD %s 0 %%%d", a, i)
	}
	out[codeSize-3] = fmt.Sprintf("ADD 0 __func_%s %%%d", name, stackSize-1)
	out[codeSize-2] = fmt.Sprintf("JZ 0 __func_%s", name)
	out[codeSize-1] = fmt.Sprintf("REL %d", stackSize)
	return out
}

func ret(name string, value int) []string {
	out := make([]string, 0)
	return out
}
