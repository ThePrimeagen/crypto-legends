package trader_test

import (
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

func TestTradeAccumulator(t *testing.T) {
    client := TestIrcClient{make(chan irc.IrcMessage, 1)}
    tlm := trader.NewTwitchLosesMoney(&client)

    tlm.Bsh.Start()
    tlm.Bsh.ReceiveMessage("foo", trader.Sell)
    tlm.Bsh.ReceiveMessage("foo2", trader.Buy)
    tlm.Bsh.ReceiveMessage("foo3", trader.Buy)
    tlm.Bsh.ReceiveMessage("foo2", trader.Sell)
    res := tlm.Bsh.GetResults()

    if res != trader.Sell {
        t.Errorf("Expected Sell but got %s", res);
    }
}
