package main

import (
	"github.com/gdamore/tcell"
	"github.com/sachaos/tview"
)

func preparePacketList() *tview.Table {
	table := tview.NewTable().
		SetBorders(false).
		SetSeparator(tview.GraphicsVertBar)
	table.SetTitle("Packets").SetBackgroundColor(tcell.ColorDefault).SetBorder(true)

	column := []string{"No.", "Time", "Flow", "Lenght", "Network", "Transport"}
}
