package handy

import (
	"net"
	"time"
)

// Session represents log data per row
type Session struct {
	start    time.Time
	durarion int64
	srcIP    net.IP
	dstIP    net.IP
	srcPort  int8
	dstPort  int8
	Protocol int8
	sendByte int64
	recvByte int64
	sendPkt  int64
	recvPkt  int64
	from     string
}

type Slot struct {
	period   Period
	sessions []Session
}
