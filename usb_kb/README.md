# USB HID Keypad Example

Defines the RGB colour and keycode for up to 16 keys.
Edit the ***keys*** array to define your own colours and keycodes.

```go
var keys = [keypad.NUM_PADS]PadInfo{
	// R, G, B, Keyscan code
	{0x00, 0xff, 0x00, keyboard.KeyUpArrow},
	...
}
```
## Building and flashing to Pico

Hold bootsel button down while plugging in to computer to enable bootloader mode.
If the pico already has tinygo firmware installed then the tinygo loader will automatically put the board 
into bootloader mode without pressing the bootsel button.

```bash
git clone git@github.com:0xcafed00d/pico_rgb_keypad.git
cd pico_rgb_keypad
cd usb_kb
tinygo flash -target pico
```
