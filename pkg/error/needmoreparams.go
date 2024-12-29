package error

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

// ERR_NEEDMOREPARAMS
// https://modern.ircdocs.horse/#errneedmoreparams-461
func NeedMoreParams(w io.Writer, srvName string, c client.Client, command string) error {
	code := 461
	nn := c.Nickname
	if nn == "" {
		nn = "*"
	}
	msg := fmt.Sprintf(":%s %d %s %s :Not enough parameters\r\n",
		srvName, code, command, nn)
	if _, err := w.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write NeedMoreParams message: %w", err)
	}
	return nil
}
