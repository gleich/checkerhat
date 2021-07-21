package main

import (
	"fmt"
	"log"

	"github.com/gleich/arcade/pkg/board"
	"github.com/gleich/lumber"
	"github.com/gliderlabs/ssh"
	"github.com/nathany/bobblehat/sense/screen"
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
		fmt.Println("Handeling ssh session")
	})
	lumber.FatalMsg(ssh.ListenAndServe(":3000", nil))

	// Flush the screen
	err = screen.Clear()
	if err != nil {
		log.Fatal(err)
	}
}
