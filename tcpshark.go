package main

import (
	"log"

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
