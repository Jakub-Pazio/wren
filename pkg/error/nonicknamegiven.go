package error

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

// ERR_NONICKNAMEGIVEN
// https://modern.ircdocs.horse/#errnonicknamegiven-431
func NoNickNameGiven(w io.Writer, srvName string, c client.Client) error {
	code := 431
	nn := c.Nickname
	if nn == "" {
		nn = "*"
	}
	msg := fmt.Sprintf(":%s %d %s :No nickname given\r\n",
		srvName, code, nn)
	if _, err := w.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write NoNicknameGiven message: %w", err)
	}
	return nil
}
