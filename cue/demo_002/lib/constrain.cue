package stdlib

// Constrain
import (
	"net"
	"time"
)

// string with ip format
ip: net.IPv4
ip: "10.0.0.1"

// string with time format
ts: time.Format(time.ANSIC)
ts: "Mon Jan 2 15:27:09 2022"