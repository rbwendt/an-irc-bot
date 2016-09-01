package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "regexp"
)

type IrcConnection struct {
	conn net.Conn
	address string
	nick string
	channel string
	reader *bufio.Reader
}

func (c *IrcConnection) Run() error {
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
		matches := r.FindStringSubmatch(msg)
		
		if len(matches) > 1 {
			c.Say(fmt.Sprintf("PONG %", matches[1]))
		}
	}
}

func (c *IrcConnection) Say(msg string) {
	fmt.Println(msg)
	fmt.Fprintf(c.conn, fmt.Sprintf("%s\n", msg))
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

func main() {
	c, err := NewIrcConnection("irc.freenode.net:6667", "IrcBot70100", "fun_channel")
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	c.Run()
}


