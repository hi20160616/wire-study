package main

import (
	"fmt"
	"os"

	"github.com/hi20160616/wire-study/tutorial"
)

// func main() {
//         message := tutorial.NewMessage()
//         greeter := tutorial.NewGreeter(message)
//         event := tutorial.NewEvent(greeter)
//
//         event.Start()
// }

func main() {
	e, err := tutorial.InitializeEvent("Good one!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e.Start()
}
