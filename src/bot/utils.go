package bot

import (
	"strings"
)

func NewLineSplit(str string) []string {
	rstr := strings.Replace(string(str), "\r", "\n", -1)
	rstr = strings.Replace(string(rstr), "\n\n", "\n", -1)
	st := strings.Split(string(rstr), "\n")
	return st
}

func PrepForIRC(str string, target string, channel string) string {
	dorp := strings.Replace(string(str), "{n}", "\x02"+target+"\x02", -1)
	dorp = strings.Replace(string(dorp), "{i}", myIP, -1)
	dorp = strings.Replace(string(dorp), "{c}", channel, -1)
	return dorp
}
