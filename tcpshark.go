package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/gdamore/tcell"
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

func NewTcpterm(src *gopacket.PacketSource, debug bool) *Tcpterm {
	view := tview.NewApplication()

	packetList := preparePacketList()
	packetDetail := preparePacketDetail()
	packetDump := preparePacketDump()

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(packetList, 0, 1, true).
		AddItem(packetDetail, 0, 1, false).
		AddItem(packetDump, 0, 1, false)
	frame := prepareFrame(layout)

	view.SetRoot(frame, true).SetFocus(packetList)

	var w io.Writer
	if debug {
		w = os.Stderr
	} else {
		w = ioutil.Discard
	}

	app := &Tcpterm{
		src:        src,
		view:       view,
		primitives: []tview.Primitive{packetList, packetDetail, packetDump},
		table:      packetList,
		detail:     packetDetail,
		dump:       packetDump,
		frame:      frame,
		logger:     log.New(w, "[tcpterm]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}
	app.SwitchToTailMode()

	view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlC {
			app.Stop()
		}

		if event.Key() == tcell.KeyTAB {
			app.rotateView()
		}
		return event
	})

	packetList.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			app.SwitchToTailMode()
		}

		if key == tcell.KeyEnter {
			app.SwitchToSelectMode()
		}
	})

	packetList.SetSelectionChangedFunc(func(row int, column int) {
		app.displayDetailOf(row)
	})

	return app
}

func (app *Tcpterm) PacketListGenerator(refreshTrigger chan bool) {
	cnt := 0
	for {
		packet, err := app.src.NextPacket()
		if err == io.EOF {
			return
		} else if err == nil {
			cnt++
			rowCount := app.table.GetRowCount()

			app.logger.Printf("count: %v start\n", cnt)

			app.table.SetCell(rowCount, 0, tview.NewTableCell(strconv.Itoa(cnt)))
			app.table.SetCell(rowCount, 1, tview.NewTableCell(packet.Metadata().Timestamp.Format(timestampFormt)))
			app.table.SetCell(rowCount, 2, tview.NewTableCell(flowOf(packet)))
			app.table.SetCell(rowCount, 3, tview.NewTableCell(strconv.Itoa(packet.Metadata().Length)))
			app.table.SetCell(rowCount, 4, tview.NewTableCell(packet.Layers()[1].LayerType().String()))
			if len(packet.Layers()) > 2 {
				app.table.SetCell(rowCount, 5, tview.NewTableCell(packet.Layers()[2].LayerType().String()))
			}

			app.packets = append(app.packets, packet)

			app.logger.Printf("count: %v end\n", cnt)
		}
	}
}
