package main

import (
	"machine"
	"machine/usb/hid/keyboard"

	keypad "github.com/0xcafed00d/pico_rgb_keypad"
)

type PadInfo struct {
	r, g, b byte
	code    keyboard.Keycode
}

var keys = [keypad.NUM_PADS]PadInfo{
	// ROW 1
	{},
	{0x00, 0xff, 0x00, keyboard.KeyUpArrow},
	{},
	{0xff, 0x50, 0x00, keyboard.KeyBackspace},
	// ROW 2
	{0x00, 0xff, 0x00, keyboard.KeyLeftArrow},
	{},
	{0x00, 0xff, 0x00, keyboard.KeyRightArrow},
	{},
	// ROW 2
	{},
	{0x00, 0xff, 0x00, keyboard.KeyDownArrow},
	{},
	{},
	// ROW 3
	{},
	{},
	{},
	{0xff, 0xff, 0x00, keyboard.KeyEnter},
}

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pad := keypad.PicoRGBKeypad{}
	pad.Init()
	buttons := keypad.ButtonState{}
	buttons.Init()

	k := keyboard.New()

	for {
		buttons.SetState(pad.GetButtonStates())

		led.Low()
		for i := 0; i < keypad.NUM_PADS; i++ {
			if buttons.IsPressed(i) && keys[i].code != 0 {
				led.High()
				pad.Illuminate(i, 0x40, 0x40, 0x40)

			} else {
				pad.Illuminate(i, keys[i].r, keys[i].g, keys[i].b)
			}

			if buttons.JustPressed(i) {
				k.Down(keys[i].code)
			}

			if buttons.JustReleased(i) {
				k.Up(keys[i].code)
			}
		}
		pad.Update()
	}
}
