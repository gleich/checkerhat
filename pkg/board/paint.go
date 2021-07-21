package board

import "github.com/nathany/bobblehat/sense/screen"

// Update the board
func updateBoard(fb *screen.FrameBuffer, grid Board) error {
	for x, row := range grid {
		for y, pixel := range row {
			fb.SetPixel(x, y, pixel)
		}
	}
	return screen.Draw(fb)
}
