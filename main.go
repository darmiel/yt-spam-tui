package main

import (
	"fmt"
	"github.com/rivo/tview"
	"time"
)

func main() {
	prim := func(text string) tview.Primitive {
		return tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(text)
	}

	logTxt := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	w := tview.ANSIWriter(logTxt)

	grid := tview.NewGrid().
		SetRows(0, 10).
		SetColumns(30, 0).
		SetBorders(true)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.
		AddItem(prim("Selected"), 0, 0, 1, 1, 0, 0, false).
		AddItem(prim("Main-View"), 0, 1, 1, 1, 0, 0, false).
		AddItem(logTxt, 1, 0, 1, 2, 0, 0, false)

	app := tview.NewApplication().SetRoot(grid, true).SetFocus(grid)

	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case d := <-t.C:
				app.QueueUpdateDraw(func() {
					_, _ = fmt.Fprintf(w, "[%s]: Hey! [red]what's up?[white]\n", d.Format("15:04:05"))
				})
			}
		}
	}()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
