package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
	"os/exec"
)

const (
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
	c []bool // Board cells
	h, w int // height and width
	l int // length of cell array
}

func BoardInit(h, w int) Board {
	b := Board { h: h, w: w, l: h * w }
	b.c = make([]bool, b.l, b.l)

	return b
}

func (b *Board) X(i int) int {
	return pymod(i, b.w)
}

func (b *Board) Y(i int) int {
	return (i - b.X(i)) / b.h
}

// get index of cell (x,y)
func (b *Board) Index(x, y int) int {
	return pymod(x, b.w) + pymod(y, b.h) * b.w
}

// set state of cell at (x,y)
func (b *Board) Set(x, y int, val bool) {
	b.c[b.Index(x, y)] = val
}

// get state of cell at (x,y)
func (b *Board) Get(x, y int) bool {
	return b.c[b.Index(x, y)]
}

func (b *Board) Neighbours(x, y int) int {
	var sum int = 0

	for offy := -1; offy <= 1; offy++ {
		for offx := -1; offx <= 1; offx++ {
			if b.Get(x + offx, y + offy) {
				sum += 1
			}
		}
	}

	return sum
}

func (b *Board) Iter() {
	bprime := BoardInit(b.h, b.w)

	// this is stupidly slow
	for i,_ := range b.c {
		x := b.X(i)
		y := b.Y(i)
		n := b.Neighbours(x, y)

		switch n {
		case 3:
			bprime.Set(x, y, true)
			break
		case 4:
			bprime.Set(x, y, b.Get(x, y))
			break
		default:
			bprime.Set(x, y, false)
			break
		}
	}

	b.c = bprime.c
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

	for i,c := range b.c {
		if b.X(i) == 0 {
			fmt.Print("║")
		}

		if c {
			fmt.Print("██")
		} else {
			fmt.Print("  ")
		}

		if b.X(i) == (b.w - 1) {
			fmt.Println("║")
		}
	}

	fmt.Print("╚")
	for x := 0; x < b.w; x++ {
		fmt.Print("══")
	}
	fmt.Println("╝")
}

/*
* credit to https://github.com/pinpox
*/
func (b *Board) Populate(percent int) {
	n := (percent * b.l) / 100

	for i := 0; i < n; i++ {
		b.c[i] = true
	}

	cells := b.c
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// shuffle shuffle shuffle
	for i := b.l; i > 0; i-- {
		random := r.Intn(i)
		cells[i - 1], cells[random] = cells[random], cells[i - 1]
	}
	b.c = cells
}

func main() {
	b := BoardInit(HEIGHT, WIDTH)
	b.Populate(50)

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
