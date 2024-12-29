package error

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

func ErroneusNickname(w io.Writer, srvName string, c client.Client, nick string) error {
	code := 432
	nn := c.Nickname
	if nn == "" {
		nn = "*"
	}
	msg := fmt.Sprintf(":%s %d %s %s :Erroneus nickname\r\n",
		srvName, code, nn, nick)
	if _, err := w.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write ErroneusNickname message %w", err)
	}
	return nil
}
