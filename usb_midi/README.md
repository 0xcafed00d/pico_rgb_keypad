
# USB MIDI drum pad Example

Each pad plays a different MIDI sound. Uses General MIDI drum sounds on channel 10.

Place the names of the drums you want each button to play into the ***sounds*** array.

```go
var sounds = [keypad.NUM_PADS]midi.Note{
	BassDrum1,
	AcousticSnare,
	ClosedHiHat,
	OpenHiHat,
	...
}
```

## Building and flashing to Pico
Hold bootsel button down while plugging in to computer to enable bootloader moder.
If the pico already has tinygo firmware installed then the tinygo loader will automatically put the board 
into bootloader mode without pressing the bootsel button.

```bash
git clone git@github.com:0xcafed00d/pico_rgb_keypad.git
cd pico_rgb_keypad
cd usb_midi
tinygo flash -target pico
```



