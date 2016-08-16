package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func main() {
	var state_names = map[uint8]string{
		1:  "established",
		2:  "syn_sent",
		3:  "syn_recv",
		4:  "fin_wait1",
		5:  "fin_wait2",
		6:  "time_wait",
		7:  "close",
		8:  "close_wait",
		9:  "last_ack",
		10: "listen",
		11: "closing",
		12: "max_states",
	}

	interval := int64(30)
	s := os.Getenv("DYNO_SOCKET_STATS_INTERVAL")
	if s != "" {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil || i < 1 {
			fmt.Println("Ignoring DYNO_SOCKET_STATS_INTERVAL value:", s)
		} else {
			interval = i
		}
	}

	source := os.Getenv("DYNO")

	for {
		time.Sleep(time.Duration(interval) * time.Second)

		statuses := make(map[uint8]int)

		sockets, _ := linuxproc.ReadNetTCPSockets("/proc/net/tcp", linuxproc.NetIPv4Decoder)
		for _, e := range sockets.Sockets {
			statuses[e.Status]++
		}

		fmt.Printf("ps=%s", source)
		for key, value := range statuses {
			fmt.Printf(" %s=%d", state_names[key], value)
		}
		fmt.Printf("\n")
	}
}
