package message

import (
	"github.com/Jakub-Pazio/wren/pkg/client"
)

type Nick struct {
	NewName client.Nickname
	OldName client.Nickname
}
