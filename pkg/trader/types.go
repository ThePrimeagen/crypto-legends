package trader

import (
	"sync"

	irc "github.com/theprimeagen/crypto-legends/pkg/irc"
)

type TradeType string

const (
	Buy  TradeType = "Buy"
	Sell           = "Sell"
	HODL           = "HodlBABY"
)

type TwitchLosesMoney struct {
	client  irc.IrcClient
	traders map[string]TradeType
    lock sync.Mutex
}
