package main

import "github.com/rivo/tview"

func main() {
	prim := func(text string) tview.Primitive {
		return tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(text)
	}

	grid := tview.NewGrid().
		SetRows(0, 10).
		SetColumns(30, 0).
		SetBorders(true)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.
		AddItem(prim("Selected"), 0, 0, 1, 1, 0, 0, false).
		AddItem(prim("Main-View"), 0, 1, 1, 1, 0, 0, false).
		AddItem(prim("Log"), 1, 0, 1, 2, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
