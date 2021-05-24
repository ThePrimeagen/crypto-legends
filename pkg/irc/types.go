package irc

type IrcMessage struct {
    Name string
    Message string
}

type LoggingCallback func(string, string)
type IrcClient interface {
    Channel() chan IrcMessage
}
