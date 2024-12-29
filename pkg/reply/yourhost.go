package reply

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

// RPL_YOURHOST
// https://modern.ircdocs.horse/#rplyourhost-002
func YourHost(w io.Writer, srvName string, c client.Client, version string) error {
	msg := fmt.Sprintf(":%s 002 %s :Your host is %s, running version %s\r\n",
		srvName, c.Nickname, srvName, version)
	_, err := w.Write([]byte(msg))
	return err
}
