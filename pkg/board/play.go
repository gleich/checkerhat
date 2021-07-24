package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

const invalidFormatMsg = "Please provide a valid coordinate in the format of: x,y"

var (
	Player1Turn = true
	Player2Turn = false
)

// Move a piece
func Play(
	s ssh.Session,
	terminal *term.Terminal,
	validPositions []Coordinate,
	playerNumber int,
) error {
	// Checking if it is the player's turn
	if (playerNumber == 1 && Player2Turn) || (playerNumber == 2 && Player1Turn) {
		// for {

		// }
	}

	// Get the proposed location
	var proposedLoc Coordinate
	for {
		var proposedLocRaw string
		fmt.Fprintln(s, &proposedLocRaw)
		responseChunks := strings.Split(strings.ToLower(strings.TrimSpace(proposedLocRaw)), ",")
		if len(responseChunks) == 2 {
			x, err := strconv.Atoi(responseChunks[0])
			if err != nil {
				fmt.Fprintln(s, invalidFormatMsg)
			}
			y, err := strconv.Atoi(responseChunks[1])
			if err != nil {
				fmt.Fprintln(s, invalidFormatMsg)
			}
			proposedLoc = Coordinate{x, y}
			break
		} else {
			fmt.Fprintln(s, invalidFormatMsg)
		}
	}

	// Checking to make sure the loc is valid
	var valid bool
	for _, validLoc := range validPositions {
		if proposedLoc == validLoc {
			valid = true
			break
		}
	}
	if !valid {
		fmt.Fprintln(s, "Please provide a valid coordinate")
	}
	return nil
}
