package main

import (
	"machine"
	"machine/usb/midi"
	"time"

	"github.com/0xcafed00d/pico_rgb_keypad/keypad"
)

func main() {
	time.Now()

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pad := &keypad.PicoRGBKeypad{}
	pad.Init()
	buttons := keypad.ButtonState{}
	buttons.Init()

	m := midi.New()

	for {

		buttons.SetState(pad.GetButtonStates())

		led.Low()
		for i := 0; i < keypad.NUM_PADS; i++ {
			if buttons.JustPressed(i) {
				led.High()
				m.NoteOn(0, 9, midi.Note(35+i), 255)
			}

			if buttons.JustReleased(i) {
				m.NoteOff(0, 9, midi.Note(35+i), 255)
			}

		}
		pad.Update()
	}
}
