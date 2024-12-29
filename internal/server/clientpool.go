package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Jakub-Pazio/wren/pkg/client"
)

type Client struct {
	Name string
	// We will allow for just one connection per user
	// for now we will just belive that each user has own id
	ConnId int
}

type clientPool struct {
	Clients map[string]Client
}

func NewClientPool() clientPool {
	return clientPool{Clients: make(map[string]Client)}
}

func (cp *clientPool) AddClientNick(name string, cid int) error {
	user := Client{Name: name, ConnId: cid}
	normalizedName := strings.ToLower(name)
	_, ok := cp.Clients[normalizedName]
	if ok {
		return errors.New("client in already in server")
	}

	cp.Clients[normalizedName] = user

	return nil
}

func (cp *clientPool) UpdateClientNick(newNick, oldNick string, cid int) error {
	oldNickNormal := strings.ToLower(oldNick)
	_, ok := cp.Clients[oldNickNormal]
	if !ok {
		return fmt.Errorf("user named %s has no active connection", oldNick)
	}
	newNickNormal := strings.ToLower(newNick)
	_, ok = cp.Clients[newNickNormal]
	if ok {
		return errors.New("client in already in server")
	}
	err := client.ValidateNickname(newNick)
	if err != nil {
		return fmt.Errorf("name %s is not valid", err)
	}
	delete(cp.Clients, oldNickNormal)
	cp.Clients[strings.ToLower(newNick)] = Client{Name: newNick, ConnId: cid}
	return nil
}

func (cp *clientPool) RemoveClient(nick string) error {
	oldNick := strings.ToLower(nick)
	delete(cp.Clients, oldNick)
	return nil
}

func (cp *clientPool) ListUsers() map[string]Client {
	return cp.Clients
}
