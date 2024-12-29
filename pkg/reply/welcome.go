package reply

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

// RPL_WELCOME
// https://modern.ircdocs.horse/#rplwelcome-001
func Welcome(w io.Writer, srvName string, c client.Client) error {
	msg := fmt.Sprintf(":%s 001 %s :Welcome to the IRC Network, %s!%s@%s\r\n",
		srvName, c.Nickname, c.Nickname, c.Realname, c.Hostname)
	_, err := w.Write([]byte(msg))
	return err
}
