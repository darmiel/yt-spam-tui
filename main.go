package main

import (
	"fmt"
	"github.com/rivo/tview"
	"strings"
	"time"
)

type Author struct {
	ID       string
	UserName string
	URL      string
}

type Comment struct {
	Author *Author
	Body   string
	Time   time.Time
}

var (
	AuthorDaniel = &Author{
		ID:       "a5Bs52",
		UserName: "Daniel",
		URL:      "https://youtube.com/c/daniel",
	}
	AuthorMax = &Author{
		ID:       "h7ab3G",
		UserName: "Max",
		URL:      "https://youtube.com/c/test123",
	}
)

var comments = []*Comment{
	{
		Author: AuthorDaniel,
		Body:   "Hello World!",
		Time:   time.Now(),
	},
	{
		Author: AuthorDaniel,
		Body:   "what's up?",
		Time:   time.Now().Add(-time.Hour).Add(-30 * time.Minute),
	},
	{
		Author: AuthorDaniel,
		Body:   "test",
		Time:   time.Now().Add(-time.Hour),
	},
	{
		Author: AuthorMax,
		Body:   "hey!",
		Time:   time.Now().Add(-24 * time.Hour),
	},
}

func main() {

	logTxt := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	grid := tview.NewGrid().
		SetRows(0, 10).
		SetColumns(30, 0).
		SetBorders(true)

	//
	commentView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	list := tview.NewList().SetHighlightFullLine(true)

	selected := func() {
		text, secondary := list.GetItemText(list.GetCurrentItem())
		_, _ = fmt.Fprintln(logTxt, "[green]Authors [white]Selected:[yellow]", text, "/", secondary)

		commentView.Clear()
		// get all comments by author
		for i, c := range comments {
			a := c.Author
			if a.URL != secondary {
				continue
			}

			headerLen := len(fmt.Sprintf("(%d) %s -- %s:", i, c.Time.Format("02.04.2006 15:04:05"), a.UserName))

			_, _ = fmt.Fprintf(commentView, "[orange](%d) [green]%s[white] -- [yellow]%s:\n\n[reset]%s[white]\n\n%s\n\n",
				i,
				c.Time.Format("02.04.2006 15:04:05"),
				a.UserName,
				c.Body,
				strings.Repeat("-", headerLen))
		}
	}

	authors := make(map[string]*Author)
	for _, a := range comments {
		authors[a.Author.ID] = a.Author
	}
	var emptyR rune
	for _, a := range authors {
		list = list.AddItem(fmt.Sprintf("%s (%s)", a.UserName, a.ID), a.URL, emptyR, selected)
	}
	//

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.
		AddItem(list, 0, 0, 1, 1, 0, 0, false).
		AddItem(commentView, 0, 1, 1, 1, 0, 0, false).
		AddItem(logTxt, 1, 0, 1, 2, 0, 0, false)

	app := tview.NewApplication().SetRoot(grid, true).SetFocus(list)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
