package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	CELLW = 4
	CELLH = 4
)

var grid [][]byte

func generation(g [][]byte) [][]byte {
	gprime := make([][]byte, len(g), len(g))
	for i,s := range g {
		gprime[i] = make([]byte, len(s), len(s))
	}

	x := len(g[0])
	y := len(g)

	// this is stupidly slow
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			var n byte
			for v := i - 1; v <= (i + 1) && (i + 1) < y; v++ {
				if v < 0 {
					continue
				}
				for h := j - 1; h <= (j + 1) && (j + 1) < x; h++ {
					if h < 0 {
						continue
					}
					n += g[v][h]
				}
			}

			n -= g[i][j]

			if n < 2 && n > 3 {
				gprime[i][j] = 0
			} else if g[i][j] == 1 && (n == 2 || n == 3) {
				gprime[i][j] = 1
			} else if g[i][j] == 0 && n == 3 {
				gprime[i][j] = 1
			}

		}
	}

	return gprime
}

func draw(g [][]byte, imd *imdraw.IMDraw) {
	x := 0.0
	y := 0.0
	for _,r := range g {
		for j,c := range r {
			if c == 0 {
				continue
			}
			imd.Color = colornames.White
			imd.Push(pixel.V(x, y), pixel.V(x + CELLW, y + CELLH))
			imd.Rectangle(0)
			x += CELLW
			y += CELLH
			if j % len(g[0]) == 0 {
				x = 0
			}
		}
	}
}

func run() {
	cfg := pixelgl.WindowConfig {
		Title: "Pixel",
		Bounds: pixel.R(0, 0, 1024, 900),
		Resizable: true,
		Undecorated: true,
		VSync: false,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	for !win.Closed() {
		win.Clear(colornames.Black)
		imd.Clear()

		grid = generation(grid)
		draw(grid, imd)

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	grid = [][]byte{
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
	}

	pixelgl.Run(run)
}
