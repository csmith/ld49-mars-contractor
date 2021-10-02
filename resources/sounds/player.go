package sounds

import (
	"bytes"
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"log"
	"math/rand"
)

const sampleRate = 44100

var (
	chill1     []byte
	chill2     []byte
	ai1        []byte
	errorBytes []byte
	beepBytes  []byte

	context = audio.NewContext(sampleRate)
)

func init() {
	var err error
	chill1, err = resources.LoadAssetBytes("music-chill1.wav")
	if err != nil {
		log.Fatal(err)
	}

	chill2, err = resources.LoadAssetBytes("music-chill2.wav")
	if err != nil {
		log.Fatal(err)
	}

	ai1, err = resources.LoadAssetBytes("music-ai1.wav")
	if err != nil {
		log.Fatal(err)
	}

	errorBytes, err = resources.LoadAssetBytes("error.wav")
	if err != nil {
		log.Fatal(err)
	}

	beepBytes, err = resources.LoadAssetBytes("beep.wav")
	if err != nil {
		log.Fatal(err)
	}
}

var backgroundBytes = make(chan []byte, 10)

type backgroundReader struct {
	next []byte
}

func (b *backgroundReader) Read(p []byte) (int, error) {
	if len(b.next) == 0 {
		b.next = <-backgroundBytes
	}

	n := len(p)
	if len(b.next) < n {
		n = len(b.next)
	}

	copy(p, b.next[:n])
	b.next = b.next[n:]
	return n, nil
}

var EnableAiBackground = false

func PlayBackground() {
	go func() {
		for {
			var track []byte
			if EnableAiBackground && rand.Intn(100) < 50 {
				track = ai1
			} else if rand.Intn(100) < 50 {
				track = chill1
			} else {
				track = chill2
			}

			src, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(track))
			if err != nil {
				panic(err)
			}
			for err == nil {
				var n int
				buf := make([]byte, sampleRate)
				n, err = src.Read(buf)
				if err == nil {
					backgroundBytes <- buf[:n]
				}
			}
		}
	}()

	player, err := audio.NewPlayer(context, &backgroundReader{})
	if err == nil {
		player.SetVolume(0.75)
		player.Play()
	}
}

func PlayError() {
	src, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(errorBytes))
	if err == nil {
		player, err := audio.NewPlayer(context, src)
		if err == nil {
			player.Play()
		}
	}
}

func PlayBeep() {
	src, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(beepBytes))
	if err == nil {
		player, err := audio.NewPlayer(context, src)
		if err == nil {
			player.Play()
		}
	}
}
