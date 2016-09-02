package irc_connection

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
)

type IrcConnection struct {
	conn net.Conn
	address string
	nick string
	channel string
	reader *bufio.Reader
}

type IrcMessageHandler interface {
	Handle(string) error
}

func (c *IrcConnection) Run(h IrcMessageHandler) error {
	r, _ := regexp.Compile("^PING :([^\n]+)\n")
	defer c.conn.Close()
	for{
		msg, err := c.reader.ReadString('\n')
		if err != nil {
			return err
		}
		if len(msg)>0 {
			fmt.Printf(msg)
		}
		err = h.Handle(msg)
		if err != nil {
			return err
		}
		
		matches := r.FindStringSubmatch(msg)
		if len(matches) > 1 {
			c.Say("PONG " + matches[1])
		}
	}
}

func (c *IrcConnection) Say(msg string) {
	fmt.Println(msg)
	fmt.Fprintf(c.conn, fmt.Sprintf("%s\n", msg))
}

func (c *IrcConnection) ChannelSay(msg string) {
	c.Say("PRIVMSG #" + c.channel + " :" + msg)
}

func NewIrcConnection(address string, nick string, channel string) (IrcConnection, error) {
	connection := IrcConnection{}
	connection.address = address
	connection.channel = channel
	connection.nick = nick
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return connection, err
	}
	connection.conn = conn
	readBuf := bufio.NewReader(connection.conn)

	connection.reader = readBuf
	connection.Say(fmt.Sprintf("NICK %s", connection.nick))
	connection.Say(fmt.Sprintf("USER ircbot 0 * %s", connection.nick))
	connection.Say(fmt.Sprintf("JOIN #%s", connection.channel))
	connection.Say(fmt.Sprintf("PRIVMSG #%s :I am here!", connection.channel))

	return connection, nil
}

