package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	trackerNumber := flag.String("tracking", "RR080368026DE", "tracking shipment status")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println(" ______________  _______                        ")
		fmt.Println(" ___  __ \\__  / / /__  /                       ")
		fmt.Println(" __  / / /_  /_/ /__  /                         ")
		fmt.Println(" _  /_/ /_  __  / _  /___                       ")
		fmt.Println(" /_____/ /_/ /_/  /_____/                       ")
		fmt.Println(" ~ DHL Tracker CLI 1.0 ~                        ")
		fmt.Println("                                                ")
		fmt.Println("Starting to track your shipment:", *trackerNumber)

		GetTrackingShipment(trackerNumber)

	} else {
		fmt.Println(" ______________  _______                          ")
		fmt.Println(" ___  __ \\__  / / /__  /                         ")
		fmt.Println(" __  / / /_  /_/ /__  /                           ")
		fmt.Println(" _  /_/ /_  __  / _  /___                         ")
		fmt.Println(" /_____/ /_/ /_/  /_____/                         ")
		fmt.Println(" ~ DHL Tracker CLI 1.0 ~                          ")
		fmt.Println("                                                  ")
		fmt.Println("Let's track your shipment, e.g. package, letter.  ")
		fmt.Println("                                                  ")
		fmt.Println("The tracking number is mandatory!                 ")
		fmt.Println("                                                  ")
		fmt.Println("Here is an example on how to use the DHL CLI:     ")
		fmt.Println("./DHL -tracking RR000000001DE                     ")
		fmt.Println("Error: only 1 tracking number per run is expected.")
		os.Exit(1)
	}
}
