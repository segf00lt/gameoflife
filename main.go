package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
)

const (
	HEIGHT = 40
	WIDTH = 50
)

var SEED [][]int = [][]int {
	{2, 1}, {3, 1},        {5, 1}, {6, 1},
		{3, 2},        {5, 2},
		{3, 3},        {5, 3},
	{2, 4}, {3, 4},        {5, 4}, {6, 4},

	{26, 1}, {3, 1},        {6, 1}, {6, 1},
		{27, 2},        {6, 2},
		{27, 3},        {6, 3},
	{26, 4}, {3, 4},        {6, 4}, {6, 4},


	{17, 3},
        {18, 1},
        {18, 3},
        {19, 2},
        {19, 3},

	{14, 3},
        {15, 1},
        {15, 3},
        {16, 2},
        {16, 3},
}

type Board struct {
	m [][]bool // Board cell matrix
	h, w int // height and width
}

func BoardInit(h, w int, seed [][]int) Board {
	b := Board { h: h, w: w }
	b.m = make([][]bool, h, h)
	for i,_ := range b.m {
		b.m[i] = make([]bool, w, w)
	}

	// seed life
	for _,pair := range seed {
		i := pair[0]
		j := pair[1]
		b.m[i][j] = true
	}

	return b
}

func (b *Board) In(i, j int) bool {
	return (i >= 0 &&
		i < b.h &&
		j >= 0 &&
		j < b.w)
}

func (b *Board) Neighbours(i, j int) int {
	offset := []int {-1, 0, 1}
	var sum int = 0

	for _,off1 := range offset {
		for _,off2 := range offset {
			if !b.In(i + off1, j + off2) {
				continue
			}
			if !(off1 == 0 && off2 == 0) && b.m[i + off1][j + off2] {
				sum += 1
			}
		}
	}

	return sum
}

func (b *Board) Iter() {
	bprime := BoardInit(b.h, b.w, nil)

	// this is stupidly slow
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			n := b.Neighbours(i, j)

			if b.m[i][j] && (n < 2 || n > 3) {
				bprime.m[i][j] = false
			} else if b.m[i][j] && (n == 2 || n == 3) {
				bprime.m[i][j] = true
			} else if !b.m[i][j] && n == 3 {
				bprime.m[i][j] = true
			}

		}
	}

	b.m = bprime.m
}

/*
* credit to https://github.com/pinpox for this Print function
*/
func (b *Board) Print() {
	fmt.Print("╔")
	for x := 1; x <= b.w; x++ {
		fmt.Print("══")
	}
	fmt.Println("╗")

	for _,r := range b.m {
		fmt.Print("║")
		for _,c := range r {
			if c {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("║")
	}

	fmt.Print("╚")
	for x := 1; x <= b.w; x++ {
		fmt.Print("══")
	}
	fmt.Println("╝")
}

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	b := BoardInit(HEIGHT, WIDTH, SEED)

	for i := 0; ; i++ {
		cmd.Run()
		b.Print()
		b.Iter()
		fmt.Println("generation:", i)
		time.Sleep(100 * time.Millisecond)
	}
}
