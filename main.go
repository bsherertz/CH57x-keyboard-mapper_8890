package main

import (
	"fmt"
	"os"

	"github.com/achushu/hid"
)

var (
	Custom = []Sequence{
		{CTRL, F1}, {CTRL, F2}, {CTRL, F3}, {CTRL, F4},
		{CTRL, F5}, {CTRL, F6}, {CTRL, F7}, {CTRL, F8},
		{CTRL, F9}, {CTRL, F10}, {CTRL, F11}, {NOMOD, KP_7},

		{NOMOD, KP_1}, {NOMOD, KP_2}, {NOMOD, KP_3},
		{NOMOD, KP_4}, {NOMOD, KP_5}, {NOMOD, KP_6},
	}
)

func main() {
	var err error

	if !hid.Supported() {
		fmt.Println("this platform / binary does not support HID")
		os.Exit(1)
	}

	dev := SelectInterface()

	if dev.Path == "" {
		fmt.Printf("could not find the device interface")
		os.Exit(2)
	}

	kbd, err := NewKeyboard(dev)
	if err != nil {
		fmt.Printf("error opening device %s\n%s\n", dev.Path, err)
		os.Exit(2)
	}
	defer kbd.Close()
	fmt.Println("connected to keyboard")

	err = kbd.SendHello()
	if err != nil {
		fmt.Println("error writing to device:", err)
	}
	fmt.Println("sent hello")

	kbd.BindMapping(MapKeys(Custom))
	fmt.Println("done!")
}

func SelectInterface() hid.DeviceInfo {
	var info hid.DeviceInfo

	devices := hid.Enumerate(VENDOR_ID, PRODUCT_ID)
	if len(devices) == 0 {
		fmt.Println("no macro keyboard detected")
		os.Exit(1)
	}
	for _, d := range devices {
		if d.Interface == INTERFACE {
			return d
		}
	}
	return info
}
