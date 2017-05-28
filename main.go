package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Starting PiCar")

	car := CreateCar()

	defer log.Println("Shutting down PiCar")
	defer car.Close()

	log.Println("starting motors...")
	car.forward()
	time.Sleep(time.Second)
	log.Println("Stopping motors...")
	car.stop()

	car.backward()

	time.Sleep(time.Second)
	car.stop()
}
