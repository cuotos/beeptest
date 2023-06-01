package main

import (
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

const (
	wordsPerMin = 20
)

var (
	tone beep.Streamer
	buf  *beep.Buffer
)

// wpm returns a duration of a Dit to satisfy the required words per minute
func wpm(wpm int) time.Duration {
	msPerDir := 60000 / (50 * wpm)
	return time.Millisecond * time.Duration(msPerDir)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	sr := beep.SampleRate(44100)
	tone, _ = generators.SinTone(sr, 700)

	speaker.Init(sr, sr.N(time.Second/10))

	buf = beep.NewBuffer(beep.Format{
		SampleRate:  sr,
		NumChannels: 1,
		Precision:   1,
	})

	// the buffer contains a dash, and dits just play a third of that, so the buffer is 3 dits long
	buf.Append(beep.Take(sr.N(wpm(wordsPerMin)*3), tone))

	done := make(chan bool)

	s := beep.Seq(
		paris(),
		paris(),
		paris(),
		paris(),
		paris(),
	)

	speaker.Play(beep.Seq(s, beep.Callback(func() { done <- true })))

	<-done
	time.Sleep(time.Second)
	return nil
}

func dot() beep.Streamer {
	return beep.Seq(
		buf.Streamer(0, buf.Len()/3),
		beep.Silence(buf.Len()/3),
	)
}

func dash() beep.Streamer {
	return beep.Seq(
		buf.Streamer(0, buf.Len()),
		beep.Silence(buf.Len()/3),
	)
}

func charSpace() beep.Streamer {
	return beep.Silence((buf.Len() / 3) * 2)
}

func wordSpace() beep.Streamer {
	return beep.Silence((buf.Len() / 3) * 6)
}

func paris() beep.Streamer {
	return beep.Seq(
		dot(), dash(), dash(), dot(), charSpace(),
		dot(), dash(), charSpace(),
		dot(), dash(), dot(), charSpace(),
		dot(), dot(), charSpace(),
		dot(), dot(), dot(),
		wordSpace(),
	)
}
