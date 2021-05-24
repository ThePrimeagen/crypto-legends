package trader_test

import (
	"log"
	"testing"

	"github.com/theprimeagen/crypto-legends/pkg/irc"
	"github.com/theprimeagen/crypto-legends/pkg/trader"
)

type TestIrcClient struct {
    channel chan irc.IrcMessage
}

func (t *TestIrcClient) Channel() chan irc.IrcMessage {
    return t.channel
}

func TestTwitchLosesMyMoney(t *testing.T) {
    client := TestIrcClient{make(chan irc.IrcMessage, 1)}
    tlm := trader.NewTwitchLosesMoney(&client)

    tlm.Start()
    client.channel <- irc.IrcMessage{Name: "foo", Message: "sell"}
    client.channel <- irc.IrcMessage{Name: "foo2", Message: "buy"}
    client.channel <- irc.IrcMessage{Name: "foo3", Message: "sell"}
    res := tlm.GetResults()

    if res != trader.Sell {
        t.Errorf("Expected Sell but got %s", res);
    }
}
