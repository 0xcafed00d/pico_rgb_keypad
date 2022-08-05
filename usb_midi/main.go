package main

import (
	"machine"

	keypad "github.com/0xcafed00d/pico_rgb_keypad/keypad"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	keyPad := &keypad.PicoRGBKeypad{}
	keyPad.Init()
	keyPad.Update()

	//m := midi.New()
	//m.NoteOn(0, 0, midi.A5, 255)

	for {
		bstates := keyPad.GetButtonStates()

		led.Low()
		for i := 0; i < keypad.NUM_PADS; i++ {
			if bstates&(1<<i) != 0 {
				led.High()
				keyPad.Illuminate(i, 0xff, 0, 0)
			} else {
				keyPad.Illuminate(i, 0, 0, 0)
			}
		}
		keyPad.Update()
	}
}
