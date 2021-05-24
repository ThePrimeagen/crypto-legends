package irc

import (
	"log"
	"os"
	"time"

	twitchIrc "github.com/gempir/go-twitch-irc"
)

type Twitch struct {
    channel chan IrcMessage
	client        *twitchIrc.Client
	enableLogging bool
	callbacks     []LoggingCallback
}

func CreateIrcClient() *Twitch {
	return &Twitch{make(chan IrcMessage, 1), nil, false, nil}
}

func (t *Twitch) Connect() error {

	token := os.Getenv("TWITCH_OAUTH_TOKEN")
	channel := os.Getenv("TWITCH_OAUTH_NAME")

	t.client = twitchIrc.NewClient(channel, token)
	t.client.Join(channel)
    t.client.OnNewMessage(func(_ string, user twitchIrc.User, message twitchIrc.Message) {
        t.channel <- IrcMessage{user.DisplayName, message.Text}
    })

	t.enableLogging = false

	go func() {
        for msg := range t.channel {
            if t.enableLogging  {
                log.Printf("%d : IRC : %s: %s\n", time.Now().Unix(), msg.Name, msg.Message)
            }
        }
	}()

	return t.client.Connect()
}

func (t *Twitch) setLogging(enabled bool) {
	t.enableLogging = enabled
}

func (t *Twitch) Channel() chan IrcMessage {
    return t.channel
}
