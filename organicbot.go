package main

import (
    "github.com/thoj/go-ircevent"
    "github.com/cdorsey/organicbot/config"
    "github.com/cdorsey/organicbot/utils"
)

const server = "irc.twitch.tv:6667"

func main() {
    config := config.Getconfig()
    irccon := irc.IRC(config.Nick, config.Nick)
    irccon.Password = config.TwitchKey

    if config.Debug {
        irccon.VerboseCallbackHandler = true
        irccon.Debug = true
    }

    irccon.AddCallback("001", func (event *irc.Event) {
        irccon.Join(config.Channel)
        irccon.Privmsg(config.Channel, "Hello World!")
    })
    irccon.AddCallback("PRIVMSG", func(event *irc.Event) {
        irccon.Quit()
    })
    err := irccon.Connect(server); utils.HandleError(err, true)
    irccon.Loop()
}