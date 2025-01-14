package main

import(
	"fmt"
	"time"
)

func main() {
	// Print "Hello World!"
	fmt.Printf("Hello world!\n")

	// Get the current date and time and print it
	currentTime := time.Now()
	fmt.Printf("It's %s now ", currentTime.Format("2006-01-02 15:04:05"))

	// Print the last 4 digits of my CWID
	cwid := 4967
	fmt.Printf("and my CWID ending is %d\n", cwid)
}