package main

import (
	"machine"
	"machine/usb/midi"
	"time"

	keypad "github.com/0xcafed00d/pico_rgb_keypad"
)

// place the names of the drums you want each button to play
// into this array
var sounds = [keypad.NUM_PADS]midi.Note{
	BassDrum1,
	AcousticSnare,
	ClosedHiHat,
	OpenHiHat,
	ElectricSnare,
	SideStick,
	CrashCymbal1,
	Tambourine,
	HandClap,
	Maracas,
	LowTom,
	HighTom,
	OpenTriangle,
	Cowbell,
	HiBongo,
	LowBongo,
}

func showPress(i int, pad *keypad.PicoRGBKeypad) {
	for r := 255; r > 0; r -= 10 {
		pad.Illuminate(i, byte(r), 0, 0)
		time.Sleep(10 * time.Millisecond)
	}
	pad.Illuminate(i, 0, 0, 0)
}

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pad := keypad.PicoRGBKeypad{}
	pad.Init()
	buttons := keypad.ButtonState{}
	buttons.Init()

	m := midi.New()

	for {
		buttons.SetState(pad.GetButtonStates())

		for i := 0; i < keypad.NUM_PADS; i++ {
			if buttons.JustPressed(i) {
				m.NoteOn(0, 9, sounds[i], 255)
				go showPress(i, &pad)
			}

			if buttons.JustReleased(i) {
				m.NoteOff(0, 9, sounds[i], 255)
			}
		}
		pad.Update()
		time.Sleep(time.Millisecond)
	}
}
