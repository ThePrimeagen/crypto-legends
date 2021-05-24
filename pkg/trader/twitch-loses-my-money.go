package trader

import (
	"log"
	"sync"

	"github.com/theprimeagen/crypto-legends/pkg/irc"
)

func NewTwitchLosesMoney(ircClient irc.IrcClient) *TwitchLosesMoney {
	t := TwitchLosesMoney{ircClient, make(map[string]TradeType), sync.Mutex{}}

	go func() {
		for msg := range ircClient.Channel() {
            t.lock.Lock()
            message := msg.Message
            name := msg.Name
            log.Printf("name: %s message: %s\n", name, message)

			switch TradeType(message) {
			case Buy:
				t.traders[name] = Buy
			case Sell:
				t.traders[name] = Sell
			case HODL:
				t.traders[name] = HODL
			default:
				// ... don't do anything ...
				// Please.  Don't do it.
			}
            t.lock.Unlock()
		}
	}()

	return &t
}

func (t *TwitchLosesMoney) Start() {
	t.traders = make(map[string]TradeType)
}

func (t *TwitchLosesMoney) GetResults() TradeType {
    t.lock.Lock()
    defer t.lock.Unlock()

	buys := 0
	sells := 0
	hodls := 0

	for _, v := range t.traders {
		switch v {
		case Buy:
			buys++
		case Sell:
			sells++
		case HODL:
			hodls++
		}
	}

	if buys > sells && buys > hodls {
		return Buy
	} else if sells > hodls {
		return Sell
	}

	return HODL
}
