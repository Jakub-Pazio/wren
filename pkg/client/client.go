package client

type Client struct {
	Nickname string
	Username string
	Hostname string
	Realname string
}

func New(nickN, userN, hostN, realN string) Client {
	return Client{nickN, userN, hostN, realN}
}
