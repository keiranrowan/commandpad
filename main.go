package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func main() {
	defer midi.CloseDriver()

	app := &cli.App{
		Name: "commandpad",
		Usage: "trigger commands from launchpad devices",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name: "Keiran Rowan",
				Email: "keiranrowan@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Value: false},
		},
		Action: poll,
	}
	app.Run(os.Args)
}

func poll(ctx *cli.Context) error {
	ParseConfig()

	in, err := midi.FindInPort(Config.General.Device)
	if err != nil {
		log.Fatal("Could not find Launchpad Input Port")
	}

	out, err := midi.FindOutPort(Config.General.Device)
	if err != nil {
		log.Fatal("Count not find Launchpad Output Port")
	}
	rd := bytes.NewReader(BuildColors())

	smf.ReadTracksFrom(rd).Do(func(ev smf.TrackEvent) {
		fmt.Printf("track %v @%vms 5s\n", ev.TrackNo, ev.AbsMicroSeconds/1000, ev.Message)
	}).Play(out)

	// color := midi.NoteOn(Config.General.Channel, midi.C(6), 64)
	// send, err := midi.SendTo(out)
	// send(color)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	stop, err := midi.ListenTo(in, func(msg midi.Message, tsms int32) {
		var bt []byte
		var ch, key, vel uint8
		switch {
		case msg.GetSysEx(&bt):
			if ctx.Bool("verbose") {
				fmt.Printf("got sysex % X\n", bt)
			}
		case msg.GetNoteStart(&ch, &key, &vel):
			if ctx.Bool("verbose") {
				fmt.Printf("starting note %s on channel %v with velocity %v\n", midi.Note(key), ch, vel)
			}
			r := reflect.ValueOf(Config.Notes)
			v := reflect.Indirect(r).FieldByName(midi.Note(key).String())
			if v.IsValid() && v.Kind() == reflect.Slice {
				entry := v.Interface().([]string)
				cmd := entry[0]
				var args []string
				if len(entry) > 1 {
					args = entry[1:]
				}
				out, _ := exec.Command(cmd, args...).Output()
				fmt.Printf("%s", out)
			}
		case msg.GetNoteEnd(&ch, &key):
			if ctx.Bool("verbose") {
				fmt.Printf("ending note %s on channel %v\n", midi.Note(key), ch)
			}
		default:
		}
	}, midi.UseSysEx())

	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	for {
		s := <-done
		switch s {
		case syscall.SIGINT:
			ParseConfig()
		case syscall.SIGTERM:
			stop()
			return nil
		default:
		}
	}
}
