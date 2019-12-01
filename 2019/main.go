package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ags131/adventofcode/2019/aoc"
)

func main() {
	start := time.Now()
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Day Missing")
		os.Exit(1)
	}
	day := 0
	fmt.Sscan(args[0], &day)
	if len(args) >= 2 && args[1] == "download" {
		err := aoc.DownloadInput(2019, day)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Downloaded input for day %d", day)
		return
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
