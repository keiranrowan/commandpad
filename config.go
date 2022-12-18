package main

import (
	"os"
	"log"

	"github.com/BurntSushi/toml"
)

type (
	config struct {
		General general
		Notes	notes
		Color	color
	}
	general struct {
		Device	 string
		Channel	 uint8
	}
	notes struct {
		Ab0 []string
		C2  []string
		E3  []string
		Ab4 []string
		C6  []string
		E7  []string
		Ab8 []string
		C10 []string
		C0  []string
		Db0 []string
		D0  []string
		Eb0 []string
		E0  []string
		F0  []string
		Gb0 []string
		G0  []string
		E1  []string
		F1  []string
		Gb1 []string
		G1  []string
		Ab1 []string
		A1  []string
		Bb1 []string
		B1  []string
		Ab2 []string
		A2  []string
		Bb2 []string
		B2  []string
		C3  []string
		Db3 []string
		D3  []string
		Eb3 []string
		C4  []string
		Db4 []string
		D4  []string
		Eb4 []string
		E4  []string
		F4  []string
		Gb4 []string
		G4  []string
		E5  []string
		F5  []string
		Gb5 []string
		G5  []string
		Ab5 []string
		A5  []string
		Bb5 []string
		B5  []string
		Ab6 []string
		A6  []string
		Bb6 []string
		B6  []string
		C7  []string
		Db7 []string
		D7  []string
		Eb7 []string
		C8  []string
		Db8 []string
		D8  []string
		Eb8 []string
		E8  []string
		F8  []string
		Gb8 []string
		G8  []string
		E9  []string
		F9  []string
		Gb9 []string
		G9  []string
		Ab9 []string
		A9  []string
		Bb9 []string
		B9  []string
	}
	color struct {
		Ab0 int
		C2  int
		E3  int
		Ab4 int
		C6  int
		E7  int
		Ab8 int
		C10 int
		C0  int
		Db0 int
		D0  int
		Eb0 int
		E0  int
		F0  int
		Gb0 int
		G0  int
		E1  int
		F1  int
		Gb1 int
		G1  int
		Ab1 int
		A1  int
		Bb1 int
		B1  int
		Ab2 int
		A2  int
		Bb2 int
		B2  int
		C3  int
		Db3 int
		D3  int
		Eb3 int
		C4  int
		Db4 int
		D4  int
		Eb4 int
		E4  int
		F4  int
		Gb4 int
		G4  int
		E5  int
		F5  int
		Gb5 int
		G5  int
		Ab5 int
		A5  int
		Bb5 int
		B5  int
		Ab6 int
		A6  int
		Bb6 int
		B6  int
		C7  int
		Db7 int
		D7  int
		Eb7 int
		C8  int
		Db8 int
		D8  int
		Eb8 int
		E8  int
		F8  int
		Gb8 int
		G8  int
		E9  int
		F9  int
		Gb9 int
		G9  int
		Ab9 int
		A9  int
		Bb9 int
		B9  int
	}
)

var Config config

func ParseConfig() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f := path + "/config.toml"
	if _, err := os.Stat(f); err != nil {
		log.Fatal(err)
	}

	_, err = toml.DecodeFile(f, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
