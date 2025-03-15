package main

import (
	"machine"
	"time"
)

// main monitors key press on keys attached on D10 and D11 pins and toggles
// LED based on one of the button press. It stops the whole program when the
// other button (red) is pressed.
func main() {
	buttonRed := machine.D10
	buttonGreen := machine.D11
	buttonRed.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonGreen.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// create an instance of Toggle to toggle LED
	toggle := new(Toggle)

	// generate global context, which, when canceled, terminates the whole program
	ctx := OnButtonPressedOnce(buttonRed)

	// obtain trigger channel that triggers every time green button is pressed
	trigger := OnButtonPressedRecurring(ctx, buttonGreen)

	// run forever loop
	for {
		select {
		case <-ctx.Done():
			time.Sleep(time.Microsecond * 100)
			println("interrupted: exiting main")
			return
		case <-trigger:
			println("button pressed")
			if !toggle.State() {
				machine.LED.High()
			} else {
				machine.LED.Low()
			}
		default:
		}
		time.Sleep(time.Millisecond * 100)
	}
}
