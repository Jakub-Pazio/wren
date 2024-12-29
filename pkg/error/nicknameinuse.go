package error

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

// ERR_NICKNAMEINUSE
// https://modern.ircdocs.horse/#errnicknameinuse-433
func NicknameInUse(w io.Writer, srvName string, c client.Client, newNick string) error {
	code := 433
	nn := c.Nickname
	if nn == "" {
		nn = "*"
	}
	msg := fmt.Sprintf(":%s %d %s %s :Nickname is already in use\r\n",
		srvName, code, nn, newNick)
	if _, err := w.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write NicknameInUse message: %s", err)
	}
	return nil
}
