package main

import (
	"fmt"
	"github.com/briandowns/spinner"
)

type Task struct {
	a                      string
	trackingSpinnerFactory func(string) *spinner.Spinner
}

func main() {
	t := &Task{a: "a"}
	task(t)
}

//TODO(cx): fix bug
func task(a *Task) {

	newTrackingSpinner := a.trackingSpinnerFactory

	tracker := newTrackingSpinner("test")
	tracker.FinalMSG = "Finish.\n"
	tracker.Start()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	tracker.Stop()

	fmt.Println("------------------")
}
