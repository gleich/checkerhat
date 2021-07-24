package main

import (
	"fmt"
	"log"

	"github.com/gleich/checkerhat/pkg/board"
	"github.com/gleich/checkerhat/pkg/out"
	"github.com/gleich/lumber"
	"github.com/gliderlabs/ssh"
	"github.com/nathany/bobblehat/sense/screen"
	"golang.org/x/term"
)

func main() {
	lumber.Info("Booted")
	lumber.ErrNilCheck = false
	fb := screen.NewFrameBuffer()
	grid := board.Board{}

	validCors, err := board.Reset(&grid, fb)
	if err != nil {
		lumber.Fatal(err, "Failed to reset the board")
	}
	fmt.Println(validCors)

	ssh.Handle(func(s ssh.Session) {
		terminal := term.NewTerminal(s, "")

		err := board.MovePiece(s, terminal, validCors, 1)
		if err != nil {
			out.Error(s, err, "Failed to move piece")
		}
	})
	lumber.FatalMsg(ssh.ListenAndServe(":3000", nil))

	// Flush the screen
	err = screen.Clear()
	if err != nil {
		log.Fatal(err)
	}
}
