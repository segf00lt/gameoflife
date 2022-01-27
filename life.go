package main

import (
	"fmt"
	"flag"
	"time"
	"math/rand"
	"os"
	"os/exec"
)

const (
	// defaults
	HEIGHT = 45
	WIDTH = 45
)

// python style modulo
func pymod(a, b int) int {
	m := a % b

	if m < 0 {
		m += b
	}

	return m
}

type Board struct {
	cur, next [][]bool
	w, h int // height and width
}

func NewBoard(w, h int) *Board {
	b := Board {
		w: w,
		h: h,
	}

	b.cur = make([][]bool, h)
	b.next = make([][]bool, h)
	for i := 0; i < h; i++ {
		b.cur[i] = make([]bool, w)
		b.next[i] = make([]bool, w)
	}

	return &b
}

// count number of live cells surrounding and including cell (i,j)
func (b *Board) Count(i, j int) int {
	offsets := [9][2]int {
		{-1,-1}, {-1, 0}, {-1, 1},
		{0, -1}, { 0, 0}, { 0, 1},
		{1, -1}, { 1, 0}, { 1, 1},
	}

	var sum int = 0
	for _,o := range offsets {
		// pymod makes boundaries periodic
		y := pymod(i + o[0], b.h)
		x := pymod(j + o[1], b.w)
		c := b.cur[y][x]
		if c {
			sum++
		}
	}

	return sum
}

func (b *Board) Iter() {
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			n := b.Count(i, j)

			switch n {
			case 3:
				b.next[i][j] = true
				break
			case 4:
				b.next[i][j] = b.cur[i][j]
				break
			default:
				b.next[i][j] = false
				break
			}
		}
	}

	tmp := b.cur
	b.cur = b.next
	b.next = tmp
}

/*
* credit to https://github.com/pinpox
*/
func (b *Board) Print() {
	fmt.Print("╔")
	for x := 0; x < b.w; x++ {
		fmt.Print("══")
	}
	fmt.Println("╗")

	for i := 0; i < b.h; i++ {
		fmt.Print("║")
		for j := 0; j < b.w; j++ {
			c := b.cur[i][j]

			if c {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("║")
	}

	fmt.Print("╚")
	for x := 0; x < b.w; x++ {
		fmt.Print("══")
	}
	fmt.Println("╝")
}

func (b *Board) Populate() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	cells := b.cur
	for i := 0; i < b.h; i++ {
		for j := 0; j < b.w; j++ {
			cells[i][j] = r.Int() % 2 == 0
		}
	}
}

func main() {
	var (
		w int
		h int
	)

	flag.IntVar(&w, "w", WIDTH, "Specify board width, default is 45")
	flag.IntVar(&h, "h", HEIGHT, "Specify board height, default is 45")
	flag.Parse()

	b := NewBoard(w, h)
	b.Populate()

	for i := 0; ; i++ {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		b.Print()
		b.Iter()
		fmt.Println("generation:", i)
		time.Sleep(200 * time.Millisecond)
	}
}
