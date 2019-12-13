package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ags131/adventofcode/2019/aoc/intcode"
	"github.com/ags131/adventofcode/2019/aoc/intcode/assembler"
	"github.com/google/subcommands"
)

type idisasmCmd struct {
	output string
}

func (*idisasmCmd) Name() string     { return "idisasm" }
func (*idisasmCmd) Synopsis() string { return "Intcode Disassembler" }
func (*idisasmCmd) Usage() string {
	return `Usage: iasm <input> 
	Flags:
		-o output
`
}

func (c *idisasmCmd) SetFlags(f *flag.FlagSet) {
	// f.IntVar(&r.day, "day", 0, "day to run, 0 for all")
	f.StringVar(&c.output, "output", "-", "file to write to (default stdout)")
}

func (c *idisasmCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	args := f.Args()
	if len(args) < 1 {
		fmt.Println("Input Missing")
		return subcommands.ExitFailure
	}
	var r io.Reader
	if args[0] == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}
		defer f.Close()
		r = f
	}
	prog := intcode.ReadInput(r)
	ret := assembler.Disassemble(prog)
	var w io.Writer
	if c.output == "-" {
		w = os.Stdout
	} else {
		f, err := os.Create(c.output)
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}
		defer f.Close()
		w = f
	}
	w.Write([]byte(strings.Join(ret, "\n")))
	return subcommands.ExitSuccess
}
