package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type MyWidget struct {
	widget.BaseWidget

	Cont      *fyne.Container
	text      *widget.Label
	statusBar *widget.Label
	b1        *widget.Button
	b2        *widget.Button

	count uint
}

func (t *MyWidget) Init() {
	t.b1 = widget.NewButton("1", func() {
		t.text.SetText("1")
		t.count++
		t.statusBar.SetText(fmt.Sprint(t.count))
	})
	t.b2 = widget.NewButton("2", func() { t.text.SetText("2") })
	t.statusBar = widget.NewLabel("status")
	bottom := fyne.NewContainerWithLayout(layout.NewCenterLayout(), t.statusBar)
	t.text = widget.NewLabelWithStyle("0", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})
	t.Cont = fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, bottom, nil, nil),
		bottom, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithRows(4),
			fyne.NewContainerWithLayout(layout.NewCenterLayout(), t.text),
			layout.NewSpacer(),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2), t.b1, t.b2),
			layout.NewSpacer(),
		))
}

func Load() *MyWidget {
	obj := &MyWidget{BaseWidget: widget.BaseWidget{}}
	obj.Init()
	return obj
}

func main() {
	f := app.New()
	w := f.NewWindow("")
	obj := Load()
	w.SetContent(obj.Cont)
	w.ShowAndRun()
}
