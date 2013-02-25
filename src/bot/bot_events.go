package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"irc"
)

var ircBot *IrcBot
var myIP string

func OnConnect(_event *irc.Event) {
	ircBot.onConnect(_event)
}

func OnKick(_event *irc.Event) {
	ircBot.onKick(_event)
}

func OnMsg(_event *irc.Event) {
	ircBot.onMsg(_event)
}

func (b *IrcBot) onConnect(_event *irc.Event) {
	resp, err := http.Get("http://icanhazip.com")
	if err != nil {
		b.myIp = "0.0.0.0"
	} else {
		dorp, _ := ioutil.ReadAll(resp.Body)
		b.myIp = string(dorp)
	}
	resp.Body.Close()

	// Identify if needed
	if b.Identify {
		b.ircConn.Privmsg(b.Nickserv, b.Command)
	}

	// Join all channels
	for i := range b.Channels {
		b.ircConn.Join(b.Channels[i])
	}
}

func (b *IrcBot) onKick(_event *irc.Event) {
	if strings.ToLower(_event.Message) != "leave" {
		if strings.ToLower(_event.Arguments[1]) == strings.ToLower(b.GetNick()) {
			b.ircConn.Join(_event.Arguments[0])
			b.SendMessage(_event.Arguments[0], "If you want me to stay out, kick me with the reason 'leave'.", "") // Yep
		}
	}
}

func (b *IrcBot) onMsg(_event *irc.Event) {
	fmt.Println("<["+_event.Arguments[0]+"]", _event.Nick+">", _event.Message)
	if strings.HasPrefix(_event.Message, b.Prefix) {

	}
}
