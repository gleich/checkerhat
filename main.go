package main

import (
	"log"

	"github.com/gleich/arcade/pkg/board"
	"github.com/gleich/arcade/pkg/out"
	"github.com/gleich/lumber"
	"github.com/gliderlabs/ssh"
	"github.com/nathany/bobblehat/sense/screen"
)

func main() {
	lumber.Info("Booted")
	lumber.ErrNilCheck = false
	fb := screen.NewFrameBuffer()
	grid := board.Board{}

	ssh.Handle(func(s ssh.Session) {
		err := board.Reset(&grid, fb)
		if err != nil {
			out.Error(s, err, "Failed to reset the board")
		}
	})
	lumber.FatalMsg(ssh.ListenAndServe(":4000", nil))

	// Flush the screen
	err := screen.Clear()
	if err != nil {
		log.Fatal(err)
	}
}
