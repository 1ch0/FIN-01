package go2

import "fmt"

var drivers = make(map[string]Driver)

type Driver interface {
	Open()
}

func Register(name string, driver Driver) {
	drivers[name] = driver
}

func DDD(name string) {
	var d Driver
	if d1, ok := drivers[name]; !ok {
		panic(fmt.Sprintf("no real driver registered"))
	} else {
		d = d1
	}
	d.Open()
}
