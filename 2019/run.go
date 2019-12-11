package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ags131/adventofcode/2019/aoc"
	"github.com/google/subcommands"
)

type runCmd struct {
}

func (*runCmd) Name() string     { return "run" }
func (*runCmd) Synopsis() string { return "Run day(s)" }
func (*runCmd) Usage() string {
	return `Usage: run day [suffix]
`
}

func (r *runCmd) SetFlags(f *flag.FlagSet) {
	// f.IntVar(&r.day, "day", 0, "day to run, 0 for all")
	// f.StringVar(&r.suffix, "suffix", "", "input suffix to append")
}

func (r *runCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	start := time.Now()
	args := f.Args()
	if len(args) < 1 {
		fmt.Println("Day Missing")
		return subcommands.ExitFailure
	}
	day := 0
	fmt.Sscan(args[0], &day)
	if len(args) >= 2 && args[1] == "download" {
		err := aoc.DownloadInput(2019, day)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Downloaded input for day %d", day)
		return subcommands.ExitSuccess
	}
	if day == 0 {
		wg := sync.WaitGroup{}
		wg.Add(25)
		for day = 1; day <= 25; day++ {
			go func(wg *sync.WaitGroup, day int) {
				defer wg.Done()
				runDay(day, "")
			}(&wg, day)
		}
		wg.Wait()
	} else {
		suffix := ""
		if len(args) >= 2 {
			suffix = args[1]
		}
		runDay(day, suffix)
	}
	elapsed := time.Now().Sub(start)
	fmt.Printf("Total Timing: %v\n", elapsed)
	return subcommands.ExitSuccess
}

func runDay(day int, suffix string) {
	start := time.Now()
	fn := days[day-1]
	input := aoc.NewInput(fmt.Sprintf("./input/day%d%s", day, suffix))
	defer input.Close()
	output := fn(input)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Day%d:\n  Part1: %v\n  Part2: %v\nTiming: %v\n", day, output.Part1, output.Part2, elapsed)
}
