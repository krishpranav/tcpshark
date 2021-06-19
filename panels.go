package main

import (
	"github.com/gdamore/tcell"
	"github.com/sachaos/tview"
)

func preparePacketList() *tview.Table {
	table := tview.NewTable().
		SetBorders(false).
		SetSeparator(tview.GraphicsVertBar)
	table.SetTitle("Packets").
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(true)

	columns := []string{"No.", "Time", "Flow", "Length", "Network", "Transport"}
	for i, column := range columns {
		table.SetCell(0, i,
			tview.NewTableCell(column).
				SetTextColor(tcell.ColorYellow).
				SetSelectable(false),
		)
	}
	table.SetFixed(1, 1)

	return table
}

func preparePacketDetail() *tview.TextView {
	text := tview.NewTextView()
	text.SetBorder(true).SetTitle("Detail").SetBackgroundColor(tcell.ColorDefault)
	return text
}
