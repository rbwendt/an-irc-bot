package say_hi

import (
	ircc "github.com/rbwendt/an-irc-bot/irc_connection"
	
	"regexp"
)

type SayHiMessageHandler struct {
	IrcConn ircc.IrcConnection
}

func NewSayHiMessageHandler(c ircc.IrcConnection) SayHiMessageHandler {
	return SayHiMessageHandler{IrcConn: c}
}

func (h *SayHiMessageHandler) Handle(msg string) error {
	saidHi, _ := regexp.MatchString("hi bot", msg)
	if saidHi {
		h.IrcConn.ChannelSay("Hi!")
	}
	return nil
}

