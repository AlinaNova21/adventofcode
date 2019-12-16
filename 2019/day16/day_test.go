package day16

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFFTp1(t *testing.T) {
	// return
	cases := []struct {
		Input  []int
		Output []int
		Phases int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{0, 1, 0, 2, 9, 4, 9, 8},
			4,
		},

		{
			[]int{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5},
			[]int{2, 4, 1, 7, 6, 1, 7, 6},
			100,
		},
		{
			[]int{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7},
			[]int{7, 3, 7, 4, 5, 4, 1, 8},
			100,
		},
		{
			[]int{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3},
			[]int{5, 2, 4, 3, 2, 1, 3, 3},
			100,
		},
	}
	for i, tc := range cases {
		dat := tc.Input
		for j := 0; j < tc.Phases; j++ {
			dat = FFT(dat)
		}
		match := true
		for j := 0; j < 8; j++ {
			match = match && dat[j] == tc.Output[j]
		}
		if !match {
			t.Errorf("Case %d didn't match: %v != %v", i, dat[:8], tc.Output)
		}
	}
}
func TestFFTp2(t *testing.T) {
	cases := []struct {
		Input  []int
		Output []int
		Phases int
	}{
		{
			[]int{0, 3, 0, 3, 6, 7, 3, 2, 5, 7, 7, 2, 1, 2, 9, 4, 4, 0, 6, 3, 4, 9, 1, 5, 6, 5, 4, 7, 4, 6, 6, 4},
			[]int{8, 4, 4, 6, 2, 0, 2, 6},
			100,
		},
		{
			[]int{0, 2, 9, 3, 5, 1, 0, 9, 6, 9, 9, 9, 4, 0, 8, 0, 7, 4, 0, 7, 5, 8, 5, 4, 4, 7, 0, 3, 4, 3, 2, 3},
			[]int{7, 8, 7, 2, 5, 2, 7, 0},
			100,
		},
		{
			[]int{0, 3, 0, 8, 1, 7, 7, 0, 8, 8, 4, 9, 2, 1, 9, 5, 9, 7, 3, 1, 1, 6, 5, 4, 4, 6, 8, 5, 0, 5, 1, 7},
			[]int{5, 3, 5, 5, 3, 7, 3, 1},
			100,
		},
	}
	for i, tc := range cases {
		data := tc.Input
		l := len(data)
		data2 := make([]int, l*10000)
		for i, v := range data {
			for j := 0; j < 10000; j++ {
				data2[i+(l*j)] = v
			}
		}
		fmt.Println(data2)
		offsetBytes := make([]byte, 7)
		for i := range offsetBytes {
			offsetBytes[i] = byte(data2[i]) + '0'
		}
		offset, _ := strconv.Atoi(string(offsetBytes))
		dat := data2
		for j := 0; j < tc.Phases; j++ {
			dat = FFT2(dat, offset)
		}
		output := dat[offset : offset+8]
		fmt.Println(output)
		match := true
		for j := 0; j < 8; j++ {
			match = match && output[j] == tc.Output[j]
		}
		if !match {
			t.Errorf("Case %d didn't match: %v != %v", i, output, tc.Output)
		}
	}
}
