package main

import (
	"machine"
	"machine/usb/hid/keyboard"

	"github.com/0xcafed00d/pico_rgb_keypad/keypad"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pad := &keypad.PicoRGBKeypad{}
	pad.Init()

	k := keyboard.New()

	buttons := keypad.ButtonState{}

	for {
		buttons.SetState(pad.GetButtonStates())

		for i := 0; i < keypad.NUM_PADS; i++ {
			if buttons.IsPressed(i) {
				pad.Illuminate(i, 0xff, 0, 0)
			} else {
				pad.Illuminate(i, 0, 0, 0)
			}

			if buttons.JustPressed(i) {
				k.Down(keyboard.KeyA + keyboard.Keycode(i))
			}

			if buttons.JustReleased(i) {
				k.Up(keyboard.KeyA + keyboard.Keycode(i))
			}
		}
		pad.Update()
	}
}
