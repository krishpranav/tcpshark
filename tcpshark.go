package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/sachaos/tview"
)

const (
	TailMode = iota
	SelectMode
)

type Tcpshark struct {
	src        *gopacket.PacketSource
	view       *tview.Application
	primitives []tview.Primitive
	table      *tview.Table
	detail     *tview.TextView
	dump       *tview.TextView
	frame      *tview.Frame
	packets    []gopacket.Packet
	mode       int

	logger *log.Logger
}

const (
	timestampFormt = "2006-01-02 15:04:05.000000"
)

func NewTcpStream(src *gopacket.PacketSource, debug bool) *Tcpshark {
	view := tview.NewApplication()

	packetList := preparePacketList()
	packetDetail := preparePacketDetail()
	packetDump := preparePacketDump()

	layout := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(packetList, 0, 1, true).AddItem(packetDetail, 0, 1, false).AddItem(packetDump, 0, 1, true)
	frame := prepareFrame(layout)

	view.SetRoot(frame, true)

	var w io.Writer
	if debug {
		w = os.Stderr
	} else {
		w = ioutil.Discard
	}
}
