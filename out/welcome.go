package out

import (
	"fmt"
	"time"

	"github.com/gliderlabs/ssh"
)

var (
	Player1Taken bool
	Player2Taken bool
)

// Output a welcome message to the user!
func Welcome(s ssh.Session) (error, string) {
	welcome_msg := ` Welcome to multiplayer space wars for the zephyr train! Before we launch you into orbit you might wanna review a few controls:
w - move up
s - move down
a - move left
d - move right
l - fire
`
	for _, char := range welcome_msg {
		fmt.Fprint(s, string(char))
		time.Sleep(time.Millisecond * 20)
	}

	var player string
	for {
		fmt.Fprint(s, "What player number do you want to be (1 or 2) ")
		fmt.Fscan(s, &player)
		if player != "1" && player != "2" {
			fmt.Fprintln(s, "\nPlease enter a valid player number")
			continue
		}

		if player == "1" {
			if Player1Taken {
				fmt.Fprintln(s, "Player 1 taken")
				fmt.Fprintln(s)
				continue
			} else {
				Player1Taken = true
				break
			}
		} else {
			if Player2Taken {
				fmt.Fprintln(s, "Player 2 taken")
				fmt.Fprintln(s)
				continue
			} else {
				Player2Taken = true
				break
			}
		}
	}

	return nil, player
}
