package audio

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

const audioDir = "assets/audio"

var op *oto.NewContextOptions
var ctx *oto.Context

func init() {
	op = &oto.NewContextOptions{}
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan
	ctx = otoCtx
}

func PlaySound(fileName string) {
	fileBytes, err := os.ReadFile(audioDir + "/" + fileName + ".mp3")
	if err != nil {
		log.Fatal(err)
	}

	// Convert the pure bytes into a reader object that can be used with the mp3 decoder
	fileBytesReader := bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	player := ctx.NewPlayer(decodedMp3)
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}
