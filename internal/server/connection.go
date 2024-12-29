package server

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"strings"
	"time"

	"github.com/Jakub-Pazio/wren/pkg/client"
	ircerr "github.com/Jakub-Pazio/wren/pkg/error"
	"github.com/Jakub-Pazio/wren/pkg/reply"
)

const (
	sName   = "127.0.0.1"
	version = "wren:0.1.1"
)

var datetime string

func init() {
	t := time.Now()

	datetime = t.UTC().Format("Mon Jan 2 15:04:05 MST 2006")
}

// each connection can have its "global state"
// it should not need any synchronization mechanizms cause
// each connection runs on its only one goroutine

// Those two valiables keep track if user is allowed to join channels
// and send messages to ther clients

// we "cache" users nickname to offload "server"
type connection struct {
	netConn net.Conn
	client  client.Client
	id      int
	ch      chan<- channelCommand
	rch     chan serverReply
	greeted bool
}

// runs for each connection from the client, it can have its own state
// that keep track of current conversation with client, such as:
// - clients nickname, address, previously send messages (for session)
// it should communicate with the "server" using channels
func (c *connection) Run() {
	c.client.Hostname = c.netConn.RemoteAddr().String()

	scr := bufio.NewScanner(c.netConn)
	defer c.userQuit()

	// for each command we want to check if its valid and make as much work
	// as it is possible, before we choose to perform some action in the "sever"
	// when we know exactly what to perform only then we send "task"
	for scr.Scan() {
		msg := scr.Text()
		if len(msg) == 0 {
			// https://modern.ircdocs.horse/#message-format
			// states: "If you encounter an empty message, silently ignore it."
			continue
		}
		msgSplit := strings.Split(msg, " ")
		command := msgSplit[0]
		args := msgSplit[1:]

		switch command {
		case "NICK":
			if err := c.handleNICK(args); err != nil {
			}
		case "USER":
			if err := c.handleUSER(args); err != nil {
			}
		case "QUIT":
			c.handleQUIT(args)
			c.netConn.Close()
			return
		default:
			slog.Warn("unknown command send from client", "command", command, "client", c.id)
		}
	}
}

func (c *connection) userQuit() {
	// send info to the server that user has left conversation
	c.handleQUIT([]string{""})
	fmt.Println("User has quit!")
}

func (c *connection) handleNICK(args []string) error {
	if len(args) == 0 || args[0] == "" {
		slog.Info("nick command with empty nickname")
		ircerr.NoNickNameGiven(c.netConn, sName, c.client)
		return nil
	}
	nick := args[0]
	err := client.ValidateNickname(nick)
	if err != nil {
		slog.Warn("user provided bad nick", "nick", nick)
		ircerr.ErroneusNickname(c.netConn, sName, c.client, nick)
		return nil
	}
	if c.client.Nickname == "" {
		argsServer := []string{nick}
		c.ch <- channelCommand{Name: "NewNICK", ConnId: c.id, Args: argsServer}
	} else {
		argsServer := []string{nick, c.client.Nickname}
		c.ch <- channelCommand{Name: "UpdateNICK", ConnId: c.id, Args: argsServer}
	}

	resp := <-c.rch
	// TODO: Later we should extract errors that core server can send to us
	// and creaet switch here were we will match on error and act based on value
	if resp.Err != nil {
		ircerr.NicknameInUse(c.netConn, sName, c.client, nick)
		return nil
	}
	// here we should get some response from the "serer"
	// and then perform action based on what was the result of our command
	c.client.Nickname = args[0]
	if c.shouldGreet() {
		c.greetUser(sName, version, datetime)
	}

	return nil
}

func (c *connection) handleUSER(args []string) error {
	if c.client.Username != "" {
		ircerr.AlreadyRegistered(c.netConn, sName, c.client)
		return nil
	}
	if len(args) < 3 {
		ircerr.NeedMoreParams(c.netConn, sName, c.client, "USER")
		return nil
	}
	username := args[0]
	realname := args[3]
	argsServer := []string{username, realname}
	c.ch <- channelCommand{Name: "USER", Args: argsServer}

	resp := <-c.rch
	if resp.Err != nil {
		fmt.Fprintf(c.netConn, "Something went wrong: %s\n", resp.Err)
		return resp.Err
	}
	c.client.Username = username
	c.client.Realname = realname
	if c.client.Nickname == "" {
		slog.Warn("USER send before NICK", "connection", c.id)
	}
	if c.shouldGreet() {
		c.greetUser(sName, version, datetime)
	}
	return nil
}

func (c *connection) handleQUIT(args []string) bool {
	var msg string
	if len(args) > 0 {
		msg = strings.Join(args, " ")
	}
	argsServer := []string{c.client.Nickname, msg}
	c.ch <- channelCommand{Name: "QUIT", Args: argsServer}
	slog.Info("User has quit server", c.client.Nickname, msg)
	return true
}

func (c *connection) shouldGreet() bool {
	return c.client.Nickname != "" && c.client.Username != "" && !c.greeted
}

func (c *connection) greetUser(sName, version, datetime string) {
	reply.Welcome(c.netConn, sName, c.client)
	reply.YourHost(c.netConn, sName, c.client, version)
	reply.Created(c.netConn, sName, c.client, datetime)
	reply.MyInfo(c.netConn, sName, c.client, version)
	c.greeted = true
}

type channelCommand struct {
	// This later should be changed to list of possible commands i want to
	// Be sending from connection to server using iota and we will have
	// one place in the repository where all those messages will be defined
	// maybe it in not as type safe as using Rust enums it seems good enough
	Name   string
	ConnId int
	Args   []string
}

// still not sure about this, maybe we could just get normal anwser and not IRC
// specific message that we want to send, but on the other hand its the IRC
// server so it probably should understand IRC on every layer
type serverReply struct {
	Err      error
	ErrCode  int
	Response string
}
