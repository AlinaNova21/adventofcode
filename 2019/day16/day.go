package day16

import (
	"io/ioutil"
	"strconv"

	"github.com/ags131/adventofcode/2019/aoc"
)

var patterns [][]int
var base = []int{0, 1, 0, -1}

func FFT(data []int) []int {
	l := len(data)
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		out := 0
		for j := 0; j < l; j++ {
			ind := ((j + 1) / (i + 1)) % 4
			if data[j] == 0 || base[ind] == 0 {
				continue
			}
			out += data[j] * base[ind]
		}
		if out < 0 {
			out *= -1
		}
		out %= 10
		ret[i] = out
	}
	return ret
}

func FFT2(data []int, offset int) []int {
	l := len(data)
	ret := make([]int, l)
	psum := 0
	for i := 0; i < l; i++ {
		ret[i] = data[i]
	}
	for i := offset; i < l; i++ {
		psum += data[i]
	}
	for i := offset; i < l; i++ {
		out := psum
		psum -= data[i]
		if out < 0 {
			out *= -1
		}
		ret[i] = out % 10
	}
	return ret
}

func genPatterns(l int) {
	l += 8
	patterns = make([][]int, l)
	base := []int{0, 1, 0, -1}
	for i := range patterns {
		pattern := make([]int, l)
		for j := 0; j < l; j++ {
			pattern[j] = base[(j/(i+1))%4]
		}
		patterns[i] = pattern[1:]
	}
}

// Run runs this day
func Run(input *aoc.Input) aoc.Output {
	bytes, _ := ioutil.ReadAll(input)
	bytes = bytes[:len(bytes)-1]
	data := make([]int, len(bytes))
	for i, v := range bytes {
		data[i] = int(v - '0')
	}
	phases := 100
	out := data
	for i := 0; i < phases; i++ {
		out = FFT(out)
	}
	part1 := out[:8]

	l := len(data)
	data2 := make([]int, l*10000)
	for i, v := range data {
		for j := 0; j < 10000; j++ {
			data2[i+(l*j)] = v
		}
	}
	offsetBytes := make([]byte, 7)
	for i := range offsetBytes {
		offsetBytes[i] = byte(data2[i]) + '0'
	}
	offset, _ := strconv.Atoi(string(offsetBytes))
	out = data2
	for i := 0; i < phases; i++ {
		// fmt.Println("Phase", i)
		out = FFT2(out, offset)
	}
	part2 := out[offset : offset+8]
	return aoc.Output{Part1: part1, Part2: part2}
}
