package out

import (
	"io"

	"github.com/gleich/lumber"
	"github.com/gliderlabs/ssh"
)

// Output an error to the user and log it to the console
func Error(s ssh.Session, err error, msg string) {
	contactMsg := " Please contact Matt Gleich about this bug."
	_, writeErr := io.WriteString(s, msg+contactMsg)
	if writeErr != nil {
		lumber.Error(writeErr, "Failed to write error msg")
	}
	lumber.Error(err, msg)
}
