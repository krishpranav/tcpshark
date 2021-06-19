package main

import (
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/urfave/cli"
)

var (
	snapshot_len int32 = 1024
	promiscuous  bool  = false
	err          error
	timeout      time.Duration = 100 * time.Millisecond
	handle       *pcap.Handle
)

func findDevice(c *cli.Context) string {
	if c.String("interface") != "" {
		return c.String("interface")
	}
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}
	return devices[0].Name
}
