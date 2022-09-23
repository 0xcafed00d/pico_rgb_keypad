// tiny go implementation of Pimoroni's C++ driver for their pico RGB keypad
// https://shop.pimoroni.com/products/pico-rgb-keypad-base
// https://github.com/pimoroni/pimoroni-pico/tree/main/libraries/pico_rgb_keypad
package pico_rgb_keypad

import (
	"machine"
	"time"
)

const (
	PIN_SDA  = machine.I2C0_SDA_PIN
	PIN_SCL  = machine.I2C0_SCL_PIN
	PIN_CS   = machine.GP17
	PIN_SCK  = machine.SPI0_SCK_PIN
	PIN_MOSI = machine.SPI0_SDO_PIN

	WIDTH              = 4
	HEIGHT             = 4
	NUM_PADS           = WIDTH * HEIGHT
	KEYPAD_ADDRESS     = 0x20
	DEFAULT_BRIGHTNESS = 0.5
)

type PicoRGBKeypad struct {
	spi_cs machine.Pin
	spi    *machine.SPI
	i2c    *machine.I2C

	buffer   [NUM_PADS*4 + 8]byte
	led_data []byte
}

func (t *PicoRGBKeypad) Init() {
	t.spi_cs = PIN_CS
	t.spi = machine.SPI0
	t.i2c = machine.I2C0
	t.led_data = t.buffer[4:]

	t.i2c.Configure(machine.I2CConfig{Frequency: 400000, SDA: PIN_SDA, SCL: PIN_SCL})

	t.spi_cs.Configure(machine.PinConfig{Mode: machine.PinOutput})
	t.spi_cs.High()
	t.spi.Configure(machine.SPIConfig{Frequency: 4 * 1024 * 1024, SCK: PIN_SCK, SDO: PIN_MOSI})

	t.SetBrightness(DEFAULT_BRIGHTNESS)
	t.Update()
}

func (t *PicoRGBKeypad) Update() {
	t.spi_cs.Low()
	t.spi.Tx(t.buffer[:], nil)
	t.spi_cs.High()
}

func (t *PicoRGBKeypad) SetBrightness(brightness float32) {
	if brightness < 0.0 || brightness > 1.0 {
		return
	}

	for i := 0; i < NUM_PADS; i++ {
		t.led_data[i*4] = 0b11100000 | byte(brightness*float32(0b11111))
	}
}

func (t *PicoRGBKeypad) Illuminate(i int, r byte, g byte, b byte) {
	if i < 0 || i >= NUM_PADS {
		return
	}
	offset := i * 4
	t.led_data[offset+1] = b
	t.led_data[offset+2] = g
	t.led_data[offset+3] = r
}

func (t *PicoRGBKeypad) IlluminateXY(x int, y int, r byte, g byte, b byte) {
	if x < 0 || x >= WIDTH || y < 0 || y >= HEIGHT {
		return
	}
	t.Illuminate(x+(y*WIDTH), r, g, b)
}

func (t *PicoRGBKeypad) Clear() {
	for i := 0; i < NUM_PADS; i++ {
		t.Illuminate(i, 0, 0, 0)
	}
}

func (t *PicoRGBKeypad) GetButtonStates() uint16 {
	buffer := [2]byte{}
	reg := [1]byte{0}

	t.i2c.Tx(KEYPAD_ADDRESS, reg[:], nil)
	t.i2c.Tx(KEYPAD_ADDRESS, nil, buffer[:])
	return ^(uint16(buffer[0]) | uint16(buffer[1])<<8)
}

// helper to keep track of which buttons have just been pressed or released since
// the last time the state was set and to debounce presses
type ButtonState struct {
	previous  [NUM_PADS]bool
	current   [NUM_PADS]bool
	debounce  [NUM_PADS]time.Duration
	last_time time.Time
}

func (t *ButtonState) Init() {
	t.last_time = time.Now()
}

func (t *ButtonState) SetState(states uint16) {
	now := time.Now()
	d := now.Sub(t.last_time)

	for i := 0; i < NUM_PADS; i++ {
		// get state of current button
		bstate := (states & (1 << i)) != 0

		// reduce debounce countdown for current button
		t.debounce[i] -= d

		// only record current button change if debounce timer has dropped below zero
		t.previous[i] = t.current[i]
		if t.debounce[i] < 0 {
			t.current[i] = bstate
		}

		// if button just released then start restart the debounce timer
		if t.JustReleased(i) {
			// 100ms is a long time to debounce, but the squishy buttons of the picorgbkeypad
			// are very bouncy - loosening the screws on the pico rgb help a bit.
			// mechanical buttons would require much shorter debounce time
			t.debounce[i] = 100 * time.Millisecond
		}
	}
	t.last_time = now
}

func (t *ButtonState) IsPressed(index int) bool {
	return t.current[index]
}

func (t *ButtonState) JustPressed(index int) bool {
	return (!t.previous[index]) && t.current[index]
}

func (t *ButtonState) JustReleased(index int) bool {
	return t.previous[index] && (!t.current[index])
}
