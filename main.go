package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		line  []byte
		fline []byte
		err   error
	)
	f, err := os.Open("./example.txt")
	check(err)

	rdr := bufio.NewReader(f)
	for err != io.EOF {
		line, _, err = rdr.ReadLine()
		if len(line) == 0 {
			break
			// continue
		}
		fline, _, err = rdr.ReadLine()
		moves := []int{}
		Moves(line, fline, moves)
	}
}

var memo = map[string]struct{}{}

func Moves(s, f grid, list []int) bool {
	if bytes.Equal(s, f) {
		return true
	}

	if _, ok := memo[s.String()]; ok {
		return false
	}

	memo[s.String()] = struct{}{}
	for i := 0; i < 8; i++ {
		s.Print()
		fmt.Println("---")
		s.Press(i)
		if Moves(s, f, list) {
			list = append([]int{i}, list...)
			return true
		}
	}
	return false
}

type grid []byte

func (g grid) Press(pos int) {
	switch pos {
	case 0:
		g[0] = flip(g[0])
		g[1] = flip(g[1])
		g[2] = flip(g[2])
		g[3] = flip(g[3])
		g[6] = flip(g[6])
	case 1:
		g[0] = flip(g[0])
		g[1] = flip(g[1])
		g[2] = flip(g[2])
		g[4] = flip(g[4])
		g[7] = flip(g[7])
	case 2:
		g[0] = flip(g[0])
		g[1] = flip(g[1])
		g[2] = flip(g[2])
		g[5] = flip(g[5])
		g[8] = flip(g[8])
	case 3:
		g[0] = flip(g[0])
		g[3] = flip(g[3])
		g[4] = flip(g[4])
		g[5] = flip(g[5])
		g[6] = flip(g[6])
	case 4:
		g[1] = flip(g[1])
		g[3] = flip(g[3])
		g[4] = flip(g[4])
		g[5] = flip(g[5])
		g[7] = flip(g[7])
	case 5:
		g[2] = flip(g[2])
		g[3] = flip(g[3])
		g[4] = flip(g[4])
		g[5] = flip(g[5])
		g[8] = flip(g[8])
	case 6:
		g[0] = flip(g[0])
		g[3] = flip(g[3])
		g[6] = flip(g[6])
		g[7] = flip(g[7])
		g[8] = flip(g[8])
	case 7:
		g[1] = flip(g[1])
		g[4] = flip(g[4])
		g[6] = flip(g[6])
		g[7] = flip(g[7])
		g[8] = flip(g[8])
	case 8:
		g[2] = flip(g[2])
		g[5] = flip(g[5])
		g[6] = flip(g[6])
		g[7] = flip(g[7])
		g[8] = flip(g[8])
	}
}

func flip(v byte) byte {
	if v == '0' {
		return '1'
	}

	return '0'
}

func (g grid) FlatPrint() {
	fmt.Println(g)
}

func (g grid) Print() {
	fmt.Println(g[0:3])
	fmt.Println(g[3:6])
	fmt.Println(g[6:])
}

func (g grid) String() string {
	return string(g)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
