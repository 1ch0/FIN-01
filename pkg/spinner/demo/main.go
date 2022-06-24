package main

import (
	"log"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.FinalMSG = "Complete!\n"
	s.Start()

	s.Prefix = "Colors: "
	s.UpdateCharSet(spinner.CharSets[83])

	if err := s.Color("yellow"); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(3 * time.Second)

	s.Stop() // Stop the spinner

	println("")
}
