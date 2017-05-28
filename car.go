package main

import (
	"log"

	rpio "github.com/stianeikeland/go-rpio"
)

var (
	// pin layout from https://pinout.xyz/pinout/pin15_gpio22
	// pin 38
	PinRightMotorsForward = rpio.Pin(20)
	// pin 40
	PinRightMotorsBackward = rpio.Pin(21)

	// pin 13
	PinLeftMotorsForward = rpio.Pin(27)
	//pin 15
	PinLeftMotorsBackward = rpio.Pin(22)
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

func (c *Car) stop() {
	log.Println("CAR: Stopping")
	PinRightMotorsForward.Low()
	PinRightMotorsBackward.Low()
	PinLeftMotorsForward.Low()
	PinLeftMotorsBackward.Low()
}
