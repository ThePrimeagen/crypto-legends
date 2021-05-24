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

type TradeAccumulator struct {
	traders map[string]TradeType
	lock    sync.Mutex
}

type TwitchLosesMoney struct {
    Bsh TradeAccumulator

	client  irc.IrcClient
}
