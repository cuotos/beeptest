package main

import (
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

var (
	tone beep.Streamer
	buf  *beep.Buffer
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	sr := beep.SampleRate(44100)
	tone, _ = generators.SineTone(sr, 700)

	speaker.Init(sr, sr.N(time.Second/10))

	buf = beep.NewBuffer(beep.Format{
		SampleRate:  sr,
		NumChannels: 1,
		Precision:   1,
	})

	buf.Append(beep.Take(sr.N(time.Second/3), tone))

	done := make(chan bool)

	s := beep.Seq(
		dot(), dot(), dot(), dot(), charSpace(),
		dot(), dot(), charSpace(),
		dot(), dash(), dot(), dot(), charSpace(),
		dot(), dash(), dot(), dot(), charSpace(),
		dash(), dash(), dash(),
		wordSpace(),
		dot(), dash(), dash(), charSpace(),
		dash(), dash(), dash(), charSpace(),
		dot(), dash(), dot(), charSpace(),
		dot(), dash(), dot(), dot(), charSpace(),
		dash(), dot(), dot(),
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
	return beep.Silence(buf.Len())
}

func wordSpace() beep.Streamer {
	return beep.Silence(buf.Len() * 6)
}
