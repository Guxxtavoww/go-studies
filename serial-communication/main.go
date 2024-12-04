package main

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	portName := "/dev/ttyUSB0" // Change this to your ESP32's serial port (e.g., COM3 on Windows)
	baudRate := 115200         // Set to the ESP32's baud rate

	config := &serial.Config{
		Name: portName,
		Baud: baudRate,
		ReadTimeout: time.Millisecond * 500, // Set read timeout
	}

	// Open the serial port
	port, err := serial.OpenPort(config)

	if err != nil {
		log.Fatalf("Failed to open serial port: %v", err)
	}

	defer port.Close()

	fmt.Printf("Connected to %s at %d baud.\n", portName, baudRate)

	// Read from the serial port in a separate goroutine
	go func() {
		scanner := bufio.NewScanner(port)

		for scanner.Scan() {
			fmt.Printf("Received: %s\n", scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Error reading from serial port: %v", err)
		}
	}()

	// Send data to the ESP32
	for {
		fmt.Print("Enter text to send (or 'exit' to quit): ")

		var input string

		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Exiting.")
			break
		}

		_, err := port.Write([]byte(input + "\n"))

		if err != nil {
			log.Printf("Error writing to serial port: %v", err)
		}
	}
}
