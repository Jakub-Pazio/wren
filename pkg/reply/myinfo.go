package reply

import (
	"fmt"
	"io"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

// RPL_MYINFO
// https://modern.ircdocs.horse/#rplmyinfo-004
func MyInfo(w io.Writer, srvName string, c client.Client, version string) error {
	userModes := "ao"
	channelModes := "mtov"
	msg := fmt.Sprintf(":%s 004 %s %s %s %s %s\r\n",
		srvName, c.Nickname, srvName, version, userModes, channelModes)
	_, err := w.Write([]byte(msg))
	return err
}
