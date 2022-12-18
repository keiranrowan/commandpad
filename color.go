package main

import (
	"bytes"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
)

func BuildColors() []byte {
	var (
		bf	bytes.Buffer
		clock = smf.MetricTicks(96)
		tr	smf.Track
	)

	tr.Add(0, smf.MetaMeter(4,4))
	tr.Add(0, smf.MetaTempo(100))
	tr.Add(0, smf.MetaInstrument("Brass"))
	tr.Add(0, midi.ProgramChange(0, gm.Instr_BrassSection.Value()))
	tr.Add(clock.Ticks4th()*4, midi.NoteOn(Config.General.Channel, midi.C(0), 64))
	//tr.Add(0, midi.NoteOff(Config.General.Channel, midi.C(0)))
	tr.Close(0)

	s := smf.New()
	s.TimeFormat = clock
	s.Add(tr)
	s.WriteTo(&bf)
	return bf.Bytes()
}
