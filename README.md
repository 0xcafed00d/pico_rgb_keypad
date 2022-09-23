
# tinygo Pico RGB keypad 

TinyGo driver for Pimoroni's Raspberry Pi Pico RGB keypad:

<https://shop.pimoroni.com/products/pico-rgb-keypad-base>
  
## Usage:
Make sure you have Go and tinyGo installed:

<https://go.dev/>

<https://tinygo.org/>


```go
// import keypad driver
import keypad "github.com/0xcafed00d/pico_rgb_keypad"

func main {
	// create a keypad drive object and initialise it.
	pad := keypad.PicoRGBKeypad{}
	pad.Init()

	// set global brightness
	pad.SetBrightness(0.2)

	// set pad 0 to RED
	pad.Illuminate(0, 0xff, 0, 0)

	// read buttons...
	pressed := pad.GetButtonStates()
}
```
The API is almost identical to Pimoroni's C++ driver API. Check out their documentation here:

<https://github.com/pimoroni/pimoroni-pico/tree/main/libraries/pico_rgb_keypad>


# Example Applications:

[USB HID Keypad](https://github.com/0xcafed00d/pico_rgb_keypad/blob/master/usb_kb/README.md)

[USB Midi Drumpad](https://github.com/0xcafed00d/pico_rgb_keypad/blob/master/usb_midi/README.md)
