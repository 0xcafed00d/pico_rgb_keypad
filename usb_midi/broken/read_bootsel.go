//go:build rp2040
// +build rp2040

package main

// https://github.com/raspberrypi/pico-examples/blob/master/picoboard/button/button.c

/*
#include <stdint.h>
#include <stdbool.h>

typedef volatile uint32_t io_rw_32;
typedef const volatile uint32_t io_ro_32;

#define	__STRING(x)	#x

#define __not_in_flash(group) __attribute__((section(".time_critical." group)))
#define __not_in_flash_func(func_name) __not_in_flash(__STRING(func_name)) func_name
#define __no_inline_not_in_flash_func(func_name) __attribute__((noinline)) __not_in_flash_func(func_name)

enum gpio_override {
    GPIO_OVERRIDE_NORMAL = 0,      ///< peripheral signal selected via \ref gpio_set_function
    GPIO_OVERRIDE_INVERT = 1,      ///< invert peripheral signal selected via \ref gpio_set_function
    GPIO_OVERRIDE_LOW = 2,         ///< drive low/disable output
    GPIO_OVERRIDE_HIGH = 3,        ///< drive high/enable output
};

static uint32_t __no_inline_not_in_flash_func(save_and_disable_interrupts) (void) {
    uint32_t status;
    __asm volatile ("mrs %0, PRIMASK" : "=r" (status)::);
    __asm volatile ("cpsid i");
    return status;
}

static void __attribute__((always_inline)) __no_inline_not_in_flash_func(restore_interrupts) (uint32_t status) {
    __asm volatile ("msr PRIMASK,%0"::"r" (status) : );
}

#define REG_ALIAS_XOR_BITS (0x1u << 12u)
#define hw_alias_check_addr(addr) ((uintptr_t)(addr))
#define hw_xor_alias_untyped(addr) ((void *)(REG_ALIAS_XOR_BITS | hw_alias_check_addr(addr)))

static void __no_inline_not_in_flash_func(hw_xor_bits) (io_rw_32 *addr, uint32_t mask) {
    *(io_rw_32 *) hw_xor_alias_untyped((volatile void *) addr) = mask;
}

static void __no_inline_not_in_flash_func(hw_write_masked) (io_rw_32 *addr, uint32_t values, uint32_t write_mask) {
    hw_xor_bits(addr, (*addr ^ values) & write_mask);
}

#define IO_QSPI_GPIO_QSPI_SS_CTRL_OEOVER_LSB 12u
#define IO_QSPI_GPIO_QSPI_SS_CTRL_OEOVER_BITS 0x00003000u
#define SIO_BASE 0xd0000000u

typedef struct {
    io_ro_32 cpuid;
    io_ro_32 gpio_in;
    io_ro_32 gpio_hi_in;
} sio_hw_t;
#define sio_hw ((sio_hw_t *)SIO_BASE)

typedef struct {
    io_ro_32 status;
    io_rw_32 ctrl;
} ioqspi_status_ctrl_hw_t;

#define NUM_QSPI_GPIOS 6u
typedef struct {
    ioqspi_status_ctrl_hw_t io[NUM_QSPI_GPIOS]; // 6
} ioqspi_hw_t;
#define IO_QSPI_BASE 0x40018000u
#define ioqspi_hw ((ioqspi_hw_t *)IO_QSPI_BASE)


bool __no_inline_not_in_flash_func(readbootsel)(){
    const uint32_t CS_PIN_INDEX = 1;

	uint32_t flags = save_and_disable_interrupts();

    // Set chip select to Hi-Z
    hw_write_masked(&ioqspi_hw->io[CS_PIN_INDEX].ctrl,
                    GPIO_OVERRIDE_LOW << IO_QSPI_GPIO_QSPI_SS_CTRL_OEOVER_LSB,
                    IO_QSPI_GPIO_QSPI_SS_CTRL_OEOVER_BITS);

    // Note we can't call into any sleep functions in flash right now
    for (volatile int i = 0; i < 1000; ++i);

	// The HI GPIO registers in SIO can observe and control the 6 QSPI pins.
    // Note the button pulls the pin *low* when pressed.
    bool button_state = !(sio_hw->gpio_hi_in & (1u << CS_PIN_INDEX));

    // Need to restore the state of chip select, else we are going to have a
    // bad time when we return to code in flash!
    hw_write_masked(&ioqspi_hw->io[CS_PIN_INDEX].ctrl,
                    GPIO_OVERRIDE_NORMAL << IO_QSPI_GPIO_QSPI_SS_CTRL_OEOVER_LSB,
                    IO_QSPI_GPIO_QSPI_SS_CTRL_OEOVER_BITS);

	restore_interrupts(flags);

	return button_state;
}

*/
import "C"

func ReadBootsel() bool {
	return bool(C.readbootsel())
}
