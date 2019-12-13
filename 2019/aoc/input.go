package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

func DownloadInput(year, day int) error {
	session := os.Getenv("COOKIE")
	if session == "" {
		return fmt.Errorf("COOKIE must be set")
	}
	client := resty.New()
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "session",
			Value: session,
		}).
		Get(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day))
	if err != nil {
		return fmt.Errorf("Error downloading input for day %d: %v", day, err)
	}
	if resp.StatusCode() == 200 {
		file, err := os.Create(fmt.Sprintf("input/day%d", day))
		if err != nil {
			return err
		}
		defer file.Close()
		file.Write(resp.Body())
	} else {
		return fmt.Errorf("Error writing input for day %d: %s", day, resp.Status())
	}
	return nil
}

// ReadIntSlice parses IntCode program input
func ReadIntSlice(input io.Reader) []int {
	ret := make([]int, 0)
	var v int
	for {
		n, _ := fmt.Fscanf(input, "%d", &v)
		if n == 0 {
			break
		}
		ret = append(ret, v)
	}
	return ret
}

func CloneIntSlice(ints []int) []int {
	ret := make([]int, len(ints))
	for i, v := range ints {
		ret[i] = v
	}
	return ret
}
