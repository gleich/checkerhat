package board

import (
	"time"

	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
)

// Reset the board
func Reset(grid *Board, fb *screen.FrameBuffer) (Coordinate, error) {
	validCors := Coordinate{}
	for y, row := range grid {
		skip := false
		if y%2 == 0 {
			skip = true
		}
		for x := range row {
			if skip {
				skip = false
				continue
			}
			if y < 3 {
				grid[x][y] = color.Red
			} else if y > 4 {
				grid[x][y] = color.Blue
			} else {
				grid[x][y] = color.White
			}
			validCors = append(validCors, [2]int{x, y})
			skip = true

			err := updateBoard(fb, *grid)
			if err != nil {
				return Coordinate{}, err
			}
			time.Sleep(time.Millisecond * 100)
		}
	}
	return validCors, nil
}
