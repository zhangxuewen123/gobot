package main

import (
	"os"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/ble"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ollie := ble.NewSpheroOllieDriver(bleAdaptor)

	work := func() {
		ollie.SetRGB(255, 0, 255)
		gobot.Every(3*time.Second, func() {
			ollie.Roll(40, uint16(gobot.Rand(360)))
		})
	}

	robot := gobot.NewRobot("ollieBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ollie},
		work,
	)

	robot.Start()
}