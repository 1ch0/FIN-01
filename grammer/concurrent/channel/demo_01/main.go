package main

var msg string
var done = make(chan bool)

func main() {
	go setup()
	<-done
	println(msg)
}

func setup() {
	msg = "do"
	done <- true

}
