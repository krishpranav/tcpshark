package main

import (
	"log"
	"time"

	"github.com/google/gopacket"
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

func createHandle(c *cli.Context) (*pcap.Handle, error) {
	fileName := c.String("read")
	if fileName != "" {
		return pcap.OpenOffline(fileName)
	} else {
		device := findDevice(c)
		return pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	}
}

func findSource(c *cli.Context) (*gopacket.PacketSource, func()) {
	handle, err := createHandle(c)
	if err != nil {
		log.Fatal(err)
	}

	if c.String("filter") != "" {
		err := handle.SetBPFFilter(c.String("filter"))
		if err != nil {
			log.Fatal(err)
		}
	}
	return gopacket.NewPacketSource(handle, handle.LinkType()), handle.Close
}
