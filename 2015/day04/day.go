package day04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/ags131/adventofcode/2015/aoc"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	var secret string
	fmt.Fscanf(input, "%s", &secret)
	part1 := 0
	part2 := 0
	acc := 1
	for {
		bytes := md5.Sum([]byte(secret + strconv.Itoa(acc)))
		str := hex.EncodeToString(bytes[:])
		cnt := 0
		for i := 0; i < 6; i++ {
			if str[i] == '0' {
				cnt++
			} else {
				break
			}
		}
		if cnt == 5 && part1 == 0 {
			part1 = acc
		}
		if cnt == 6 {
			part2 = acc
			break
		}
		acc++
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
