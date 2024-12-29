package error

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

func NotRegistered(w io.Writer, srvName string, c client.Client) error {
	code := 451
	nn := c.Nickname
	if nn == "" {
		nn = "*"
	}
	msg := fmt.Sprintf(":%s %d %s :You have not registered\r\n",
		srvName, code, nn)
	if _, err := w.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write ErroneusNickname message %w", err)
	}
	return nil
}
