package main

import (
	"encoding/hex"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"layeh.com/gopus"
)

func voiceToString(voiceMsg chan *discordgo.Packet) {
	var err error
	speakers := make(map[uint32]*gopus.Decoder)
	for vMsg := range voiceMsg {
		if _, ok := speakers[vMsg.SSRC]; !ok {
			speakers[vMsg.SSRC], err = gopus.NewDecoder(16000, 1)
			if err != nil {
				panic(err)
			}
		}

		convertedMsg, err := speakers[vMsg.SSRC].Decode(vMsg.Opus, 320, false)
		if err != nil {
			panic(err)
		}

		buf := make([]byte, 2*len(convertedMsg))

		for i := 0; i < len(convertedMsg); i++ {
			var h, l uint8 = uint8(i >> 8), uint8(i & 0xff)
			buf[i] = h
			buf[i+1] = l
		}

		myString := hex.EncodeToString(buf)

		fmt.Println("buf: ", myString)
	}

	fmt.Println("voice has been canceled")
}
