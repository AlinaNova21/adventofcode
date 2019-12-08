package day08

import (
	"fmt"
	"io/ioutil"

	"github.com/ags131/adventofcode/2019/aoc"
)

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	bytes, _ := ioutil.ReadAll(input)
	str := string(bytes[:len(bytes)-1])
	width, height := 25, 6
	layerLen := width * height
	layerCnt := len(str) / layerLen
	layers := make([]map[rune]int, layerCnt)
	leastZeros := 0
	for i := 0; i < layerCnt; i++ {
		cnts := make(map[rune]int, 0)
		for _, c := range str[i*layerLen : (i+1)*layerLen] {
			if _, ok := cnts[c]; ok {
				cnts[c]++
			} else {
				cnts[c] = 1
			}
		}
		layers[i] = cnts
		if cnts['0'] < layers[leastZeros]['0'] {
			leastZeros = i
		}
	}
	img := make([]rune, layerLen)
	for i := layerCnt - 1; i >= 0; i-- {
		layer := str[i*layerLen : (i+1)*layerLen]
		for j, c := range layer {
			if c == '2' {
				continue
			}
			img[j] = c
		}
	}
	part1 := layers[leastZeros]['1'] * layers[leastZeros]['2']
	part2 := "\n"
	i := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			clr := 40 + (7 * (img[i] - 48))
			part2 += fmt.Sprintf("\x1b[%dm ", clr)
			i++
		}
		part2 += fmt.Sprint("\x1b[0m\n")
	}
	return aoc.Output{Part1: part1, Part2: part2}
}
