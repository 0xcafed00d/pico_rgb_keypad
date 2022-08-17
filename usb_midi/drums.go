package main

import (
	"machine/usb/midi"
)

const (
	// general midi drum map for channel 10
	AcousticBassDrum midi.Note = iota + 35
	BassDrum1
	SideStick
	AcousticSnare
	HandClap
	ElectricSnare
	LowFloorTom
	ClosedHiHat
	HighFloorTom
	PedalHiHat
	LowTom
	OpenHiHat
	LowMidTom
	HiMidTom
	CrashCymbal1
	HighTom
	RideCymbal1
	ChineseCymbal
	RideBell
	Tambourine
	SplashCymbal
	Cowbell
	CrashCymbal2
	Vibraslap
	RideCymbal2
	HiBongo
	LowBongo
	MuteHiConga
	OpenHiConga
	LowConga
	HighTimbale
	LowTimbale
	HighAgogo
	LowAgogo
	Cabasa
	Maracas
	ShortWhistle
	LongWhistle
	ShortGuiro
	LongGuiro
	Claves
	HiWoodBlock
	LowWoodBlock
	MuteCuica
	OpenCuica
	MuteTriangle
	OpenTriangle
)
