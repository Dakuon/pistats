package main

import (
	"time"
	"github.com/goiot/devices/monochromeoled"
	"golang.org/x/exp/io/i2c"
)


func main() {
	display, err := monochromeoled.Open(&i2c.Devfs{Dev: "/dev/i2c-1"})
	if err != nil {
		panic(err)
	}
	defer display.Close()

	if err := display.Clear(); err != nil {
		panic(err)
	}

	for {
		data := getData()

		image := buildImage(data)

		if err := display.SetImage(0, 0, image); err != nil {
			panic(err)
		}

		if err := display.Draw(); err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
	}
}