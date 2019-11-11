package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	temperature,  humidity := 11.5, 80.00 

	for {
		// print temperature and humidity
		fmt.Printf("Temperature = %.2f *C, Humidity = %.2f \n",temperature,  humidity)
		time.Sleep(time.Second * 3)
	}
	
}
