package board

import (
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
)

// Reset the board
func Reset(grid *Board, fb *screen.FrameBuffer) error {
	for x, row := range grid {
		var skip bool
		for y := range row {
			if skip == true {
				skip = false
				continue
			}
			grid[x][y] = color.White
		}
	}
	return updateBoard(fb, *grid)
}
