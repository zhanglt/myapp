package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type myButton struct {
	widget.Button
}

func (m *myButton) TappedSecondary(*fyne.PointEvent) {
	log.Println("Right Click")
}

func newMyButton(label string, tapped func()) *myButton {
	ret := &myButton{}
	ret.ExtendBaseWidget(ret)
	ret.Text = label
	ret.OnTapped = tapped

	return ret
}

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		newMyButton("Right tap me", func() {
			log.Println("Normal callback")
		}),
	))

	w.ShowAndRun()
}
