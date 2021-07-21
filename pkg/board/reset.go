package board

import (
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
)

// Reset the board
func Reset(grid *Board, fb *screen.FrameBuffer) error {
	for y, row := range grid {
		skip := false
		if y%2 != 0 {
			skip = true
		}
		for x := range row {
			if skip {
				skip = false
				grid[y][x] = color.Black
				continue
			}
			grid[y][x] = color.White
			skip = true
		}
	}
	return updateBoard(fb, *grid)
}
