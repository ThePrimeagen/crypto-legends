package trader

import (
	"log"

	"github.com/theprimeagen/crypto-legends/pkg/irc"
)

func NewTwitchLosesMoney(ircClient irc.IrcClient) *TwitchLosesMoney {
	t := TwitchLosesMoney{
        TradeAccumulator{},
		ircClient,
	}

	go func() {
		for msg := range ircClient.Channel() {
            message := TradeType(msg.Message)

            if message != Buy && message != Sell && message != HODL {
                continue
            }

            t.Bsh.ReceiveMessage(msg.Name, message)
		}
	}()

	return &t
}

func (t *TradeAccumulator) ReceiveMessage(name string, message TradeType) {
    t.lock.Lock()

    switch TradeType(message) {
    case Buy:
        t.traders[name] = Buy
    case Sell:
        t.traders[name] = Sell
    case HODL:
        t.traders[name] = HODL
    default:
    }
    t.lock.Unlock()
}

func (t *TradeAccumulator) Start() {
    t.traders = make(map[string]TradeType)
}

func (t *TradeAccumulator) GetResults() TradeType {
	t.lock.Lock()
	defer t.lock.Unlock()

	buys := 0
	sells := 0
	hodls := 0

    log.Printf("Traders %d \n", len(t.traders))
	for n, v := range t.traders {
        log.Printf("Traders %s %s\n", n, v)
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
