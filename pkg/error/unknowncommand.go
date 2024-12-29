package error

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

func UnknownCommand(w io.Writer, srvName string, c client.Client, command string) error {
	nn := c.Nickname
	if nn == "" {
		nn = "*"
	}
	msg := fmt.Sprintf(":%s 421 %s %s :UnknownCommand",
		srvName, nn, command)
	if _, err := w.Write([]byte(msg)); err != nil {
		return fmt.Errorf("failed to write UnknownCommand: %w", err)
	}
	return nil
}
