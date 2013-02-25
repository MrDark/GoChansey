package bot

import (
	"encoding/json"
	"io/ioutil"

	"irc"
)

var lastcmd = int64(0)

type IrcBot struct {
	Channels []string
	Server   string
	Prefix   string
	Nick     string
	Realname string
	Cooldown int64
	Identify bool
	Nickserv string
	Command  string

	ircConn *irc.Connection
	myIp    string
}

func NewIrcBot(_configFile string) (*IrcBot, error) {
	ircBot = &IrcBot{}

	c, err := ioutil.ReadFile(_configFile)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(c), ircBot); err != nil {
		return nil, err
	}

	return ircBot, nil
}

func (b *IrcBot) Start() {
	b.ircConn = irc.IRC(b.Nick, b.Realname)
	b.ircConn.Connect(b.Server)
	b.ircConn.AddCallback("001", OnConnect)
	b.ircConn.AddCallback("PRIVMSG", OnMsg)
	b.ircConn.AddCallback("KICK", OnKick)
	b.ircConn.Loop()
}

func (b *IrcBot) SendMessage(channel string, text string, target string) {
	if len(text) == 0 {
		return
	}
	msgs := NewLineSplit(text)

	if len(msgs) >= 3 {
		msgs = msgs[0:2]
	}

	for i := range msgs {
		b.ircConn.Privmsg(channel, PrepForIRC(msgs[i], target, channel))
	}
}

func (b *IrcBot) GetNick() string {
	return b.ircConn.GetNick()
}
