package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	car *Car
)

func main() {
	log.Println("Starting PiCar")

	car := CreateCar()
	//car := CreateCarDummy()

	defer log.Println("Shutting down PiCar")
	defer car.Close()

	http.HandleFunc("/move/", handleRequest)
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Called ")
	fmt.Println(r.URL)
	direction := strings.TrimPrefix(r.URL.Path, "/move/")
	fmt.Println("Translated to " + direction)
	switch direction {
	case "forward":
		car.forward()
		fmt.Fprintf(w, "forward")
	case "backward":
		car.backward()
		fmt.Fprintf(w, "backward")
	case "left":
		car.left()
		fmt.Fprintf(w, "left")
	case "right":
		car.right()
		fmt.Fprintf(w, "right")
	case "stop":
		car.stop()
		fmt.Fprintf(w, "stop")
	default:
		car.stop()
		fmt.Fprintf(w, "default: stop")
	}
}

func CreateCarDummy() DummyCar {
	return DummyCar{}
}

type DummyCar struct {
}

func (*DummyCar) Close() {
	log.Println("DummyCar: Destroy called")
}

func (c *DummyCar) forward() {
	log.Println("DummyCar: Running forward")
}

func (c *DummyCar) backward() {
	log.Println("DummyCar: Running backward")
}

func (c *DummyCar) left() {
	log.Println("DummyCar: Turning left")
}

func (c *DummyCar) right() {
	log.Println("DummyCar: Turning right")
}

func (c *DummyCar) stop() {
	log.Println("DummyCar: Stopping")
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!doctype html>

<html lang="en">
<head>
    <meta charset="utf-8">
    <title>PiCar</title>

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
    <link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">
    <meta name="viewport" content="width=device-width,user-scalable=no">

    <script type="application/javascript">
        $(document).ready(function() {
           console.log("Document.ready - binding button handlers")
            function bind_start_and_stop_to_button(button_id, direction) {
                $("#" + button_id).on("touchstart mousedown", function (event) {
                    console.log(direction + " - mouse down (" + event.type + ")");
                    $.get("/move/" + direction);
                });
                $("#" + button_id).on("touchend mouseup", function (event) {
                    console.log(direction + " - mouse up (" + event.type + ")");
                    $.get("/move/stop");
                });
            }
            bind_start_and_stop_to_button("button_forward", "forward");
            bind_start_and_stop_to_button("button_left", "left");
            bind_start_and_stop_to_button("button_right", "right");
            bind_start_and_stop_to_button("button_backward", "backward");
            $("#button_stop").click(function() {
                console.log("Calling stop");
                $.get("/move/stop");
            });
        });
    </script>

    <style type="text/css">
        .nopadding {
            padding: 0;
        !important;
            margin: 0;
        !important;
        }
    </style>
</head>
<body>
<h1>PiCar</h1>

<div>
    <div class="row">
        <div class="col-xs-4 nopadding"></div>
        <div class="col-xs-4 nopadding">

            <button id="button_forward" type="button" class="btn btn-lg btn-primary btn-block"><i class="fa fa-arrow-up"
                                                                              aria-hidden="true"></i>
            </button>
        </div>
    </div>
    <div class="row">
        <div class="col-xs-4 nopadding">
            <button id="button_left" type="button" class="btn btn-lg btn-primary btn-block"><i class="fa fa-arrow-left"
                                                                              aria-hidden="true"></i>

            </button>
        </div>
        <div class="col-xs-4 nopadding">
            <button id="button_stop" type="button" class="btn btn-lg btn-danger btn-block"><i class="fa fa-stop" aria-hidden="true"></i>
            </button>
        </div>
        <div class="col-xs-4 nopadding">
            <button id="button_right" type="button" class="btn btn-lg btn-primary btn-block"><i class="fa fa-arrow-right"
                                                                              aria-hidden="true"></i>
            </button>
        </div>
    </div>
    <div class="row">
        <div class="col-xs-4 nopadding"></div>
        <div class="col-xs-4 nopadding">
            <button id="button_backward" type="button" class="btn btn-lg btn-primary btn-block"><i class="fa fa-arrow-down"
                                                                              aria-hidden="true"></i>
            </button>
        </div>
    </div>
</div>


<form>

</form>

</body>
</html>`)
}
