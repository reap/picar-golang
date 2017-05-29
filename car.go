package main

import (
	"log"

	rpio "github.com/stianeikeland/go-rpio"
)

var (
	// pin layout from https://pinout.xyz/pinout/pin15_gpio22
	// pin 40
	PinRightMotorsForward = rpio.Pin(21)
	// pin 38
	PinRightMotorsBackward = rpio.Pin(20)

	// pin 15
	PinLeftMotorsForward = rpio.Pin(22)
	//pin 13
	PinLeftMotorsBackward = rpio.Pin(27)
)

func CreateCar() Car {
	if err := rpio.Open(); err != nil {
		log.Fatalln(err)
	}

	PinRightMotorsForward.Output()
	PinLeftMotorsForward.Output()
	PinRightMotorsBackward.Output()
	PinLeftMotorsBackward.Output()

	var car Car
	car = Car{}
	return car
}

type Car struct {
}

func (*Car) Close() {
	rpio.Close()
}

func (c *Car) forward() {
	log.Println("CAR: Running forward")
	PinRightMotorsBackward.Low()
	PinLeftMotorsBackward.Low()
	PinRightMotorsForward.High()
	PinLeftMotorsForward.High()
}

func (c *Car) backward() {
	log.Println("CAR: Running backward")
	PinRightMotorsForward.Low()
	PinLeftMotorsForward.Low()
	PinRightMotorsBackward.High()
	PinLeftMotorsBackward.High()
}

func (c *Car) left() {
	log.Println("CAR: Turning left")
	PinRightMotorsForward.High()
	PinLeftMotorsForward.Low()
	PinRightMotorsBackward.Low()
	PinLeftMotorsBackward.High()
}

func (c *Car) right() {
	log.Println("CAR: Turning right")
	PinRightMotorsForward.Low()
	PinLeftMotorsForward.High()
	PinRightMotorsBackward.High()
	PinLeftMotorsBackward.Low()
}

func (c *Car) stop() {
	log.Println("CAR: Stopping")
	PinRightMotorsForward.Low()
	PinRightMotorsBackward.Low()
	PinLeftMotorsForward.Low()
	PinLeftMotorsBackward.Low()
}
