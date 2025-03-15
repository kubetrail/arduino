package main

import (
	"context"
	"machine"
	"time"
)

// Decoded codes corresponding to various remote keys
const (
	RemotePower    = 0x2
	RemoteSource   = 0x1
	RemoteKey1     = 0x4
	RemoteKey2     = 0x5
	RemoteKey3     = 0x6
	RemoteKey4     = 0x8
	RemoteKey5     = 0x9
	RemoteKey6     = 0xA
	RemoteKey7     = 0xC
	RemoteKey8     = 0xD
	RemoteKey9     = 0xE
	RemoteKey0     = 0x11
	RemoteDash     = 0x23
	RemotePreCh    = 0x13
	RemoteMute     = 0xF
	RemoteChList   = 0x6B
	RemoteVolUp    = 0x7
	RemoteVolDn    = 0xB
	RemoteChUp     = 0x12
	RemoteChDown   = 0x10
	RemoteMenu     = 0x1A
	RemoteSmartHub = 0x79
	RemoteGuide    = 0x4F
	RemoteTools    = 0x4B
	RemoteInfo     = 0x1F
	RemoteUp       = 0x60
	RemoteDown     = 0x61
	RemoteLeft     = 0x65
	RemoteRight    = 0x62
	RemoteEnter    = 0x68
	RemoteReturn   = 0x58
	RemoteExit     = 0x2D
	RemoteA        = 0x6C
	RemoteB        = 0x14
	RemoteC        = 0x15
	RemoteD        = 0x16
	RemoteEManual  = 0x3F
	RemoteSports   = 0xB8
	RemoteCC       = 0x25
	RemoteStop     = 0x46
	RemoteReverse  = 0x45
	RemotePlay     = 0x47
	RemotePause    = 0x4A
	RemoteForward  = 0x48
)

// RemoteKeyMap maps key codes to the names
var RemoteKeyMap = map[uint16]string{
	RemotePower:    "Power",
	RemoteSource:   "Source",
	RemoteKey1:     "1",
	RemoteKey2:     "2",
	RemoteKey3:     "3",
	RemoteKey4:     "4",
	RemoteKey5:     "5",
	RemoteKey6:     "6",
	RemoteKey7:     "7",
	RemoteKey8:     "8",
	RemoteKey9:     "9",
	RemoteKey0:     "0",
	RemoteDash:     "Dash",
	RemotePreCh:    "PreCh",
	RemoteMute:     "Mute",
	RemoteChList:   "ChList",
	RemoteVolUp:    "VolUp",
	RemoteVolDn:    "VolDn",
	RemoteChUp:     "ChUp",
	RemoteChDown:   "ChDown",
	RemoteMenu:     "Menu",
	RemoteSmartHub: "SmartHub",
	RemoteGuide:    "Guide",
	RemoteTools:    "Tools",
	RemoteInfo:     "Info",
	RemoteUp:       "Up",
	RemoteDown:     "Down",
	RemoteLeft:     "Left",
	RemoteRight:    "Right",
	RemoteEnter:    "Enter",
	RemoteReturn:   "Return",
	RemoteExit:     "Exit",
	RemoteA:        "A",
	RemoteB:        "B",
	RemoteC:        "C",
	RemoteD:        "D",
	RemoteEManual:  "EManual",
	RemoteSports:   "Sports",
	RemoteCC:       "CC",
	RemoteStop:     "Stop",
	RemoteReverse:  "Reverse",
	RemotePlay:     "Play",
	RemotePause:    "Pause",
	RemoteForward:  "Forward",
}

// Trigger defines communication channel that triggers every time
// a button is pressed
type Trigger chan struct{}

// Toggle toggles every time it is called via it's State() method
type Toggle struct {
	state bool
}

// State returns current state and flips it
func (t *Toggle) State() bool {
	state := t.state
	t.state = !t.state
	return state
}

// OnButtonPressedOnce cancels the returned context on button press.
// This is supposed to be a one time use function instance
func OnButtonPressedOnce(pin machine.Pin) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		for {
			if !pin.Get() {
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	return ctx
}

// OnButtonPressedRecurring provides event trigger on the returned channel
// every time button is pressed. The monitoring continues as long as the input
// context is valid. The returned channel is closed when the context is done.
func OnButtonPressedRecurring(ctx context.Context, pin machine.Pin) Trigger {
	trigger := make(Trigger)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(trigger)
				println("interrupted: exiting OnButtonPressedRecurring")
				return
			default:
				if !pin.Get() {
					trigger <- struct{}{}
				}
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	return trigger
}
