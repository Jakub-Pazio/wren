package reply

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

// RPL_CREATED
// https://modern.ircdocs.horse/#rplcreated-003
func Created(w io.Writer, srvName string, c client.Client, datetime string) error {
	msg := fmt.Sprintf(":%s 003 %s :This server was created %s\r\n",
		srvName, c.Nickname, datetime)
	_, err := w.Write([]byte(msg))
	return err
}
