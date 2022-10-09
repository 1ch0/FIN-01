package proxy_pattern

import "fmt"

// StationProxy 代理了 Station，代理类中持有被代理类对象，
// 并且和被代理类对象实现了同一接口。
type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点：%s，剩余: %d \n", name, station.stock)
	} else {
		fmt.Println("售空")
	}
}

type StationProxy struct {
	station *Station
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("代理点：%s，剩余: %d \n", name, proxy.station.stock)
	} else {
		fmt.Println("售空")
	}
}
