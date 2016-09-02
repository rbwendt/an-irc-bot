package main

import (
	ircc "github.com/rbwendt/an-irc-bot/irc_connection"
	sh "github.com/rbwendt/an-irc-bot/say_hi"
    
	"fmt"
	"os"
)

func main() {
	c, err := ircc.NewIrcConnection("irc.freenode.net:6667", "IrcBot70100", "fun_channel")
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	h := sh.NewSayHiMessageHandler(c)
	c.Run(&h)
}


