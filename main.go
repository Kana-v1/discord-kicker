package main

import "github.com/bwmarrin/discordgo"

var Token string = "MTAxNjQ1Mzc2MTY1MDEzMTAwNA.GP6U1m.-uHkFTUaJwvGy2Zb-Z3jNrH3vEAJ6YykiTtDK4"
var ChannelID string = "1016454712393998349"
var ServerID string = "1016454711953588334"

func main() {
	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		panic(err)
	}

	defer s.Close()

	err = s.Open()
	if err != nil {
		panic(err)
	}

	voiceStream, err := s.ChannelVoiceJoin(ServerID, ChannelID, true, false)
	if err != nil {
		panic(err)
	}

	voiceToString(voiceStream.OpusRecv)
}
