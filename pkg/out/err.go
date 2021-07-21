package out

import (
	"fmt"

	"github.com/gleich/lumber"
	"github.com/gliderlabs/ssh"
)

// Output an error to the user and log it to the console
func Error(s ssh.Session, err error, msg string) {
	fmt.Fprintln(s, msg+" Please contact Matt Gleich about this bug.")
	lumber.Error(err, msg)
}
