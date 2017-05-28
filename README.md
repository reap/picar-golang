### Rasperry Pi controlled 4WD-car

Uses GO-RPio-library to control GPIO ports (github.com/stianeikeland/go-rpio)

To install in Windows7 use

    GOOS=linux go get github.com/stianeikeland/go-rpio

To build for Rasperry Pi 2b running Raspian use:

    GOARM=7 GOARCH=arm GOOS=linux go build